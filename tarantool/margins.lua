local module = {}

report = require('report')

function module.calculate()
    local margins = {}
    margins['margin'] = 0
    margins['margin level'] = 0
    margins['margin requirement'] = 4
    return margins
end

function module.show()
    local margins = module.calculate()
    report.space()
    report.title('Margins')
    report.comment('to be implemented...')

    for name, value in pairs(margins) do
        if name == 'margin' then
            report.variable(name, value)
        else
            report.variable(name, value, '%')
        end
    end
end

return module
