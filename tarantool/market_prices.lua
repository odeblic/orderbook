local module = {}

report = require('report')

function module.init()
    box.schema.space.create('market_prices', {if_not_exists = true})
    box.space.market_prices:create_index('primary', {type = 'hash', parts = {1, 'string'}})
end

function module.update(symbol, value)
    box.space.market_prices:replace{symbol, value}
end

function module.read(symbol)
    local tuple = box.space.market_prices:get(symbol)
    return tuple and tuple[2] or nil
end

function module.list()
    local market_prices = {}
    local result = box.space.market_prices:select{}

    for _, entry in pairs(result) do
        local symbol = entry[1]
        local price = entry[2]
        table.insert(market_prices, {symbol, price})
    end

    return market_prices
end

function module.show()
    report.space()
    report.title("Market prices")

    if box.space.market_prices:len() > 0 then
        local entries = box.space.market_prices:select{}
        for _, entry in ipairs(entries) do
            local symbol = entry[1]
            local price = entry[2]
            table.insert(trade, 1, id)
            report.price(price, symbol)
        end
    else
        report.comment('<none>')
    end
end

return module
