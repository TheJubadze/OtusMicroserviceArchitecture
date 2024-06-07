$username = "root"
$password = "secret"
$database = "root"
$psqlhost = "localhost"
$port = "5432"
$scriptFile = "users.sql"

$env:PGPASSWORD = $password
psql -U $username -d $database -h $psqlhost -p $port -f $scriptFile