local module = {}

report = require('report')

function module.calculate()
    local positions = {}
    local trades = box.space.trades:select{}

    if #trades == 0 then
        return nil
    end

    for _, entry in ipairs(trades) do
        local index = entry[1]
        local trade = entry[2]
        local symbol = trade[1]
        local direction = trade[2]
        local quantity = trade[4]

        if positions[symbol] == nil then
            positions[symbol] = 0
        end

        if direction == "buy" then
            positions[symbol] = positions[symbol] + quantity
        elseif direction == "sell" then
            positions[symbol] = positions[symbol] - quantity
        end
    end

    return positions
end

function module.show()
    report.space()
    report.title('Positions')
    local positions = module.calculate()

    if positions ~= nil then
        for symbol, amount in pairs(positions) do
            report.position(amount, symbol)
        end
    else
        report.comment('<none>')
    end
end

return module
