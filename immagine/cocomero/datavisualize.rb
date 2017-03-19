#! /usr/bin/ruby

require 'gruff'
require 'json'

# Doc here:
# http://www.rubydoc.info/github/topfunky/gruff/Gruff/Bezier

class Failures_show
  def initialize()
    @failures = []
  end
  def drawbar(output)
    g = Gruff::Bar.new
    g.title = 'Failaures of Spacewalk-testsuite'
    # todo, puts maybe the date as label tagged
    # g.labels = { 0 => '5/6', 1 => '5/15', 2 => '5/24', 3 => '5/30', 4 => '6/4'}
    g.data('Failures', @failures[1])
    g.write(output)
  end
  # read array in json
  def get_results(file_json)
    json = File.read(file_json)
    @failures= JSON.parse(json)
  end
end

bezier = Failures_show.new
bezier.get_results('results.json')
bezier.drawbar("results.png")
