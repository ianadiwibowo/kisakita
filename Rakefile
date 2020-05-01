ENV['APP_ENV'] ||= 'development'

require 'dotenv'
require 'standalone_migrations'

Dotenv.load if ENV['APP_ENV'] == 'development'

StandaloneMigrations::Tasks.load_tasks
