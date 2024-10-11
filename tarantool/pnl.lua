local module = {}

report = require('report')

function module.calculate()
    local pnl = {}
    local prices = {}
    local positions = {}
    local notionals = {}
    local unrealized = {}

    pnl['initial deposit'] = 100
    pnl['equity'] = 0
    pnl['unrealized PnL'] = 0
    pnl['total notional'] = 0
    pnl['equity'] = 0

    local market_prices = box.space.market_prices:select{}

    for _, entry in pairs(market_prices) do
        local symbol = entry[1]
        local price = entry[2]
        prices[symbol] = price
    end

    local trades = box.space.trades:select{}

    for _, entry in ipairs(trades) do
        local trade = entry[2]
        local symbol = trade[1]
        local direction = trade[2]
        local price = trade[3]
        local quantity = trade[4]

        if positions[symbol] == nil then
            positions[symbol] = 0
        end

        if unrealized[symbol] == nil then
            unrealized[symbol] = 0
        end

        if direction == "buy" then
            positions[symbol] = positions[symbol] + quantity
            unrealized[symbol] = unrealized[symbol] + quantity * (prices[symbol] - price)
        elseif direction == "sell" then
            positions[symbol] = positions[symbol] - quantity
            unrealized[symbol] = unrealized[symbol] + quantity * (price - prices[symbol])
        end
    end

    for symbol, amount in pairs(unrealized) do
        local label = 'unrealized ' .. symbol
        pnl[label] = amount
    end

    for symbol, amount in pairs(positions) do
        notionals[symbol] = amount * prices[symbol]
    end

    for symbol, amount in pairs(notionals) do
        local label = 'notional ' .. symbol
        pnl[label] = amount
        pnl['total notional'] = pnl['total notional'] + amount
    end

    return pnl
end

function module.show()
    local pnl = module.calculate()
    report.space()
    report.title('PnL')
    report.comment('to be implemented...')

    for name, value in pairs(pnl) do
        report.variable(name, value)
    end
end

return module
