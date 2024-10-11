local module = {}

report = require('report')

function module.init()
    box.schema.space.create('order_queue', {if_not_exists = true})
    box.schema.sequence.create('order_id_generator')
    box.space.order_queue:create_index('primary', {type = 'tree', parts = {1, 'unsigned'}})
end

function module.push(symbol, direction, price, quantity)
    local id = box.sequence.order_id_generator:next()
    local order = {symbol, direction, price, quantity}
    box.space.order_queue:insert{id, order}
    return id
end

function module.pop()
    local result = box.space.order_queue:select({}, {limit = 1})

    if #result == 0 then
        return nil
    end

    entry = result[1]
    local id = entry[1]
    local order = entry[2]
    box.space.order_queue:delete(id)
    table.insert(order, 1, id)
    return order
end

function module.count()
    return box.space.order_queue:len()
end

function module.show()
    report.space()
    report.title('Order queue')

    if box.space.order_queue:len() > 0 then
        local result = box.space.order_queue:select{}

        for _, entry in ipairs(result) do
            local index = entry[1]
            local order = entry[2]
            report.order(index, order[1], order[2], order[3], order[4])
        end
    else
        report.comment('<none>')
    end
end

return module
