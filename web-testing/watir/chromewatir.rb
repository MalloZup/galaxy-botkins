#! /usr/bin/ruby

require 'watir'
# https://watir.com/examples/

def click_each(links)
  links.each do |link|
    puts link.href
    link.click if link.present?
  end
end

def get_href
  links.each do |link|
    puts link.href
  end
end

# init for chrome
browser = Watir::Browser.new :chrome
browser.goto 'google.com'

# wait that something is there
until browser.div(:id=>"hplogo").exists? do sleep 1 end
# fill in by id
browser.text_field(:id=> 'lst-ib').set 'watir'
browser.button(type: 'submit').click
### SCREENSHOTS

# save an entire page
browser.goto 'https://github.com/MalloZup'
browser.screenshot.save ("test.png")
######## 

### LINKS
# transform each link of webpage in an object array
links = browser.links.to_a
# experimental
click_each(links)

browser.close
