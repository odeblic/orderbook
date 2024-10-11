local module = {}

report = require('report')

function module.init()
    box.schema.space.create('order_book', {if_not_exists = true})
    box.space.order_book:create_index('primary', {type = 'hash', parts = {1, 'unsigned'}})
end

function module.show()
    report.space()
    report.title('Order book')
    report.comment('to be implemented...')
end

return module
