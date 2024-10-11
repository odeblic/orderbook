local module = {}

BLACK = 0
RED = 1
GREEN = 2
YELLOW = 3
BLUE = 4
MAGENTA = 5
CYAN = 6
WHITE = 7

function module.space()
    io.write('\n')
end

function module.line(color)
    character = character or '-'
    length = length or 50
    color = color or YELLOW
    text = string.rep(character, length)
    text = string.format("\x1b[3%dm%s\x1b[0m", color, text)
    print(text)
end

function module.title(text)
    text = string.format("[\x1b[3%dm%s\x1b[0m]", CYAN, text)
    print(text)
end

function module.message(text)
    text = string.format(">  \"\x1b[3%dm%s\x1b[0m\"", MAGENTA, text)
    print(text)
end

function module.comment(text)
    text = string.format("\x1b[3%dm%s\x1b[0m", WHITE, text)
    print(text)
end

function module.variable(name, value, unit)
    if unit == nil then
        text = string.format("%-20s\t: \x1b[3%dm%s\x1b[0m", name, MAGENTA, value)
    elseif unit == '$' then
        text = string.format("%-20s\t: \x1b[3%dm$%s\x1b[0m", name, MAGENTA, value)
    elseif unit == '%' then
        text = string.format("%-20s\t: \x1b[3%dm%s%%\x1b[0m", name, MAGENTA, value)
    end
    print(text)
end

function module.order(id, symbol, direction, price, quantity)
    direction = string.format("\x1b[3%dm%s\x1b[0m", direction_color(direction), direction)
    price = string.format("\x1b[3%dm%14.4f\x1b[0m", YELLOW, price)
    quantity = string.format("\x1b[3%dm%14.4f\x1b[0m", YELLOW, quantity)
    text = string.format("#%d\t%s\t%s\t%s  %s", id, symbol, direction, price, quantity)
    print(text)
end

function module.trade(id, symbol, direction, price, quantity)
    direction = string.format("\x1b[3%dm%s\x1b[0m", direction_color(direction), direction)
    price = string.format("\x1b[3%dm%14.4f\x1b[0m", YELLOW, price)
    quantity = string.format("\x1b[3%dm%14.4f\x1b[0m", YELLOW, quantity)
    text = string.format("#%d\t%s\t%s\t%s  %s", id, symbol, direction, price, quantity)
    print(text)
end

function module.price(price, base_currency, quote_currency)
    quote_currency = quote_currency or 'USD'
    text = string.format("1 %s = \x1b[3%dm%.4f\x1b[0m %s", base_currency, YELLOW, price, quote_currency)
    print(text)
end

function module.position(amount, currency)
    text = string.format("%s:\t\x1b[3%dm%12.4f\x1b[0m", currency, YELLOW, amount)
    print(text)
end

function make_style(fgcolor, bgcolor, format)
    local sequence = '\27['
    if fgcolor == nil and bgcolor == nil and format == nil then
        sequence = sequence .. '0'
    else
        if fgcolor ~= nil then
            sequence = sequence .. '0'
        end
    end
    sequence = sequence .. 'm'
    return sequence
end

function direction_color(direction)
    if direction == 'buy' then
        return GREEN
    elseif direction == 'sell' then
        return RED
    else
        return WHITE
    end
end

return module
