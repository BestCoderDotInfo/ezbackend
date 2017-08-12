require 'rubygems'
require 'listen'
require 'daemons'

APP_ROOT = File.expand_path(File.dirname(__FILE__))

loop do
  listener = Listen.to("#{APP_ROOT}/_posts/") do |modified, added, removed|
    exec("(cd #{APP_ROOT} && jekyll build --watch)")
  end
  listener.start
  sleep(5)
end
