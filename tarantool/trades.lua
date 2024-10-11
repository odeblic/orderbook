local module = {}

report = require('report')

function module.init()
    box.schema.space.create('trades', {if_not_exists = true})
    box.schema.sequence.create('trade_id_generator')
    box.space.trades:create_index('primary', {type = 'tree', parts = {1, 'unsigned'}})
end

function module.book(symbol, direction, price, quantity)
    local id = box.sequence.trade_id_generator:next()
    local order = {symbol, direction, price, quantity}
    box.space.trades:insert{id, order}
    return id
end

function module.cancel(id)
    box.space.trades:delete(id)
end

function module.get(id)
    local tuple = box.space.trades:get(id)
    return tuple and tuple[2] or nil
end

function module.list()
    local trades = {}
    local result = box.space.trades:select{}

    for _, entry in pairs(result) do
        local id = entry[1]
        local trade = entry[2]
        table.insert(trade, 1, id)
        table.insert(trades, trade)
    end

    return trades
end

function module.count()
    return box.space.trades:len()
end

function module.show()
    report.space()
    report.title('Trades')

    if box.space.trades:len() > 0 then
        local result = box.space.trades:select{}

        for _, entry in ipairs(result) do
            local index = entry[1]
            local trade = entry[2]
            report.trade(index, trade[1], trade[2], trade[3], trade[4])
        end
    else
        report.comment('<none>')
    end
end

return module
