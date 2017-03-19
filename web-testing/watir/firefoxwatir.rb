#! /usr/bin/ruby

require 'watir'
browser = Watir::Browser.new :firefox

browser.goto 'google.com'
until browser.div(:id=>"hplogo").exists? do sleep 1 end
browser.text_field(:id=> 'lst-ib').set 'watir'
browser.button(type: 'submit').click

puts browser.title
# => 'Hello World! - Google Search'
#browser.close
