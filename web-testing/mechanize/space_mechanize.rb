#! /usr/bin/ruby

require 'rubygems'
require 'mechanize'

@server_url = 'headref-suma3pg.mgr.suse.de'

# create Mechanize instance
@agent = Mechanize.new
@agent.verify_mode = OpenSSL::SSL::VERIFY_NONE

# initialize space server and login
def init_server(space_server_url, user, pwd)
  page = @agent.get("https://#{space_server_url}/rhn/Login.do")
  login = page.form('loginForm')
  login.username  = user
  login.password = pwd
  button = login.button_with(:value => "Sign In")
  page = @agent.submit(login, button)
end

def find_links(page)
  page.links.each do |link|
    puts link
  end
end

def main
  page = init_server(@server_url, 'admin', 'admin')
  puts page.links_with :href => /rhn/
  puts page.body
  #find_links(page)
end

main
