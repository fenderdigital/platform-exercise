I have taken course from ardanlabs, for writing service in go.
So for this task. I have taken this codebase https://github.com/ardanlabs/service basline

GO to path /app
> go run admin/main.go genkey
private and public key files generated


Build Docker image
Go to path /
> make api

To run
> make run

To migrate
> make migrate

To Seed
> make seed


Since seed add a admin user with userid, based on the token is generated using
> go run app/admin/main.go --db-disable-tls=1 gentoken 9c223318-deb4-416a-9778-fad0df4edf98 app/private.pem  RS256


Note : Build is not done, since I need to upgrade my local go version from 1.13 to .16 and since it took more than 4 hours. I am stopping here

