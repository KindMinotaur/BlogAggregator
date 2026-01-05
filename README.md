# BlogAggregator
Learning to build an RSS feed

This is a CLI tool called gator that requires a Postgres database running and Go installed on the device to run. To install gator run go install github.com/KindMinotaur/BlogAggregator. 
The tool requires a config file named .gatorconfig.json in your home folder containing:
{
  "db_url": "postgres://user:password@localhost:5432/gator?sslmode=disable"
}

Replace user, password, and gator with your own Postgres username, password, and database name.

Use commands for the tool with the syntax BlogAggregator <command> <arguments> such as: BlogAggregator register username or BlogAggregator follow <url> 
