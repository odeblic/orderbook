box.cfg {
    listen = 3301,
    memtx_memory = 104857600,
    snap_io_rate_limit = 0
}

box.schema.user.create('moneyboy', { password = 'secret1234' })
box.schema.user.grant('moneyboy', 'read,write,execute', 'universe')

report = require('report')
market_prices = require('market_prices')
order_queue = require('order_queue')
order_book = require('order_book')
trades = require('trades')
positions = require('positions')
pnl = require('pnl')
margins = require('margins')

report.space()
report.message('Creating the Tarantool database...')

market_prices.init()
order_queue.init()
order_book.init()
trades.init()

market_prices.show()
order_queue.show()
order_book.show()
trades.show()
positions.show()
pnl.show()
margins.show()

report.space()
report.message('Trading is now possible on my cross-margin matching engine!')

function status()
    market_prices.show()
    order_queue.show()
    order_book.show()
    trades.show()
    positions.show()
    pnl.show()
    margins.show()
end

function exit()
    os.exit()
end
