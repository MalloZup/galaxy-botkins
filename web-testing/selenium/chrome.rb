#! /usr/bin/ruby

require "selenium-webdriver"

# doc 
# http://www.rubydoc.info/gems/selenium-webdriver/0.0.28/Selenium/WebDriver/Element

class Chrome
  attr_reader :links_path
  def initialize(server)
    @browser = Selenium::WebDriver.for :chrome
    @browser.get server
    @links_path = Array.new
  end
  def visit
     return 0
  end	  
  def login(admin)
	  user = @browser.find_element(:id, "username-field")
	  user.send_keys admin
	  pwd = @browser.find_element(:id, "password-field")
	  pwd.send_keys admin
	  @browser.find_element(id: 'login').click
  end

  # Find for the current page all link, store the to array.
  def find_link_page
    @browser.find_elements(:tag_name, "a").each do |link|
	    #puts link.attribute("href")
	    @links_path.push(link.attribute("href"))
    end 
  end
  def navigate(url)
     @browser.navigate.to url
  end 

end

client = Chrome.new("http://headref-suma3pg.mgr.suse.de/")
client.login("admin")
client.find_link_page
# puts client.links_path
