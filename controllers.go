package main

func webFrameWorkListController() error {
	displayer("Pick a web-framework package:", "gin-gonic/gin", "gofiber/Fiber", "labstack/echo",
		"beego/beego", "Menu")
	val, err := picker()
	if err != nil {
		return err
	}

	switch val {
	case 1:
		err := runner("get", "-v", "github.com/gin-gonic/gin")
		if err != nil {
			return err
		}
	case 2:
		err := runner("get", "-v", "github.com/gofiber/fiber/v3")
		if err != nil {
			return err
		}
	case 3:
		err := runner("get", "-v", "github.com/labstack/echo/v4")
		if err != nil {
			return err
		}
	case 4:
		err := runner("get", "-v", "github.com/beego/beego/v2@latest")
		if err != nil {
			return err
		}
	case 5:
		driver()
	default:
		return errNonePicked
	}
	return nil
}

func routersListController() error {
	displayer("Pick a router package:", "gorilla/mux", "chi", "httprouter", "Menu")
	val, err := picker()
	if err != nil {
		return err
	}

	switch val {
	case 1:
		err := runner("get", "-v", "github.com/gorilla/mux")
		if err != nil {
			return err
		}
	case 2:
		err := runner("get", "-v", "github.com/go-chi/chi/v5")
		if err != nil {
			return err
		}
	case 3:
		err := runner("get", "-v", "github.com/julienschmidt/httprouter")
		if err != nil {
			return err
		}
	case 4:
		driver()
	default:
		return errNonePicked
	}
	return nil
}

func databasesListController() error {
	displayer("Pick a database package:", " pgx", " pq", " go-sql-driver/mysql",
		" go-sqlite3", " go-mssqldb", " ORacle[godror]", " couchbase/gocb",
		" mongo-driver/mongo", " Apache Cassandra[gocql]", "Menu")
	val, err := picker()
	if err != nil {
		return err
	}

	switch val {
	case 1:
		err := runner("get", "-v", "github.com/jackc/pgx/v5")
		if err != nil {
			return err
		}
	case 2:
		err := runner("get", "-v", "github.com/lib/pq")
		if err != nil {
			return err
		}
	case 3:
		err := runner("get", "-v", "github.com/go-sql-driver/mysql")
		if err != nil {
			return err
		}

	case 4:
		err := runner("get", "-v", "github.com/mattn/go-sqlite3")
		if err != nil {
			return err
		}
	case 5:
		err := runner("get", "-v", "github.com/denisenkom/go-mssqldb")
		if err != nil {
			return err
		}
	case 6:
		err := runner("get", "-v", "github.com/godror/godror")
		if err != nil {
			return err
		}
	case 7:
		err := runner("get", "-v", "github.com/couchbase/gocb/v2@latest")
		if err != nil {
			return err
		}
	case 8:
		err := runner("get", "-v", "go.mongodb.org/mongo-driver/mongo")
		if err != nil {
			return err
		}
	case 9:
		err := runner("get", "-v", "github.com/gocql/gocql")
		if err != nil {
			return err
		}
	case 10:
		driver()
	default:
		return errNonePicked
	}
	return nil
}

func configurationsListController() error {
	displayer("Pick a configuration package:", "viper", "godotenv", "goyaml", "Menu")
	val, err := picker()
	if err != nil {
		return err
	}

	switch val {
	case 1:
		err := runner("get", "-v", "github.com/spf13/viper")
		if err != nil {
			return err
		}
	case 2:
		err := runner("get", "-v", "github.com/joho/godotenv")
		if err != nil {
			return err
		}
	case 3:
		err := runner("get", "-v", "gopkg.in/yaml.v3")
		if err != nil {
			return err
		}
	case 4:
		driver()
	default:
		return errNonePicked
	}
	return nil
}

func cachesListController() error {
	displayer("Pick a cache package:", "redis", "memcached", "bigcache", "Menu")
	val, err := picker()
	if err != nil {
		return err
	}

	switch val {
	case 1:
		err := runner("get", "-v", "github.com/redis/go-redis/v9")
		if err != nil {
			return err
		}
	case 2:
		err := runner("get", "-v", "github.com/bradfitz/gomemcache/memcache")
		if err != nil {
			return err
		}
	case 3:
		err := runner("get", "-v", "github.com/allegro/bigcache/v3")
		if err != nil {
			return err
		}
	case 4:
		driver()
	default:
		return errNonePicked
	}
	return nil
}

func protocolsListController() error {
	displayer("Pick a protocol package:", " websocker", " webrtc", " paho/MQTT", " Apache/kafka",
		" Socker.IO", " rabitMQ/amqp", " NATS", " Fasthttp", " grpc-go", "Menu")
	val, err := picker()
	if err != nil {
		return err
	}

	switch val {
	case 1:
		err := runner("get", "-v", "github.com/gorilla/websocket")
		if err != nil {
			return err
		}
	case 2:
		err := runner("get", "-v", "github.com/pion/webrtc/v3")
		if err != nil {
			return err
		}
	case 3:
		err := runner("get", "-v", "github.com/eclipse/paho.mqtt.golang")
		if err != nil {
			return err
		}
	case 4:
		err := runner("get", "-v", "github.com/confluentinc/confluent-kafka-go/v2/kafka")
		if err != nil {
			return err
		}
	case 5:
		err := runner("get", "-v", "github.com/googollee/go-socket.io")
		if err != nil {
			return err
		}
	case 6:
		err := runner("get", "-v", "github.com/rabbitmq/amqp091-go")
		if err != nil {
			return err
		}
	case 7:
		err := runner("get", "-v", "github.com/nats-io/nats.go")
		if err != nil {
			return err
		}
	case 8:
		err := runner("get", "-v", "github.com/valyala/fasthttp")
		if err != nil {
			return err
		}
	case 9:
		err := runner("get", "-v", "google.golang.org/grpc")
		if err != nil {
			return err
		}
	case 10:
		driver()
	default:
		return errNonePicked
	}
	return nil
}

func testsListController() error {
	displayer("Pick a testing package:", "testify", "ginkgo", "httpexpect", "Menu")
	val, err := picker()
	if err != nil {
		return err
	}

	switch val {
	case 1:
		err := runner("get", "-v", "github.com/stretchr/testify")
		if err != nil {
			return err
		}
	case 2:
		err := runner("get", "-v", "github.com/onsi/ginkgo/v2")
		if err != nil {
			return err
		}
	case 3:
		err := runner("get", "-v", "github.com/gavv/httpexpect/v2")
		if err != nil {
			return err
		}
	case 4:
		driver()
	default:
		return errNonePicked
	}
	return nil
}

func logsListController() error {
	displayer("Pick a log package:", "zap", "logrus", "zerolog", "Menu")
	val, err := picker()
	if err != nil {
		return err
	}

	switch val {
	case 1:
		err := runner("get", "-v", "go.uber.org/zap")
		if err != nil {
			return err
		}
	case 2:
		err := runner("get", "-v", "github.com/sirupsen/logrus")
		if err != nil {
			return err
		}
	case 3:
		err := runner("get", "-v", "github.com/rs/zerolog/log")
		if err != nil {
			return err
		}
	case 4:
		driver()

	default:
		return errNonePicked
	}
	return nil
}

func cliListController() error {
	displayer("Pick a cli package:", "spf13/cobra", "urfave/cli", "spf13/pflag", "Menu")
	val, err := picker()
	if err != nil {
		return err
	}

	switch val {
	case 1:
		err := runner("get", "-v", "github.com/spf13/cobra@latest")
		if err != nil {
			return err
		}
	case 2:
		err := runner("get", "-v", "github.com/urfave/cli/v2")
		if err != nil {
			return err
		}
	case 3:
		err := runner("get", "-v", "github.com/spf13/pflag")
		if err != nil {
			return err
		}
	case 4:
		driver()
	default:
		return errNonePicked
	}
	return nil
}

func thirdPartyListController() error {
	displayer("Pick a thirdparty package:", "stripe/strip-go", "slack-go/slack", "google/go-github", "Menu")
	val, err := picker()
	if err != nil {
		return err
	}

	switch val {
	case 1:
		err := runner("get", "-v", "github.com/stripe/stripe-go/v78")
		if err != nil {
			return err
		}
	case 2:
		err := runner("get", "-v", "github.com/slack-go/slack")
		if err != nil {
			return err
		}
	case 3:
		err := runner("get", "-v", "github.com/google/go-github/v62")
		if err != nil {
			return err
		}
	case 4:
		driver()
	default:
		return errNonePicked
	}
	return nil
}

func otherListController() error {
	displayer("Utilities: Pick a package:", " ulule/Limiter", " rs/CORS", " go-fault/Fault Injection",
		" go-playground/validator", " google/uuid", " rs/xid", " cosmtrek/air", " gocolly/colly", " microcosm-cc/bluemonday",
		"graphql-go/graphql", "arl/statsviz", "gorm.io/gorm", "uber/fx", "unidoc/unioffice", "wneessen/go-mail",
		"sqlc-dev/sqlc", "Menu")
	val, err := picker()
	if err != nil {
		return err
	}

	switch val {
	case 1:
		err := runner("get", "-v", "github.com/ulule/limiter/v3@v3.11.2")
		if err != nil {
			return err
		}
	case 2:
		err := runner("get", "-v", "github.com/rs/cors")
		if err != nil {
			return err
		}
	case 3:
		err := runner("get", "-v", "github.com/lingrino/go-fault")
		if err != nil {
			return err
		}
	case 4:
		err := runner("get", "-v", "github.com/go-playground/validator/v10")
		if err != nil {
			return err
		}
	case 5:
		err := runner("get", "-v", "github.com/google/uuid")
		if err != nil {
			return err
		}
	case 6:
		err := runner("get", "-v", "github.com/rs/xid")
		if err != nil {
			return err
		}
	case 7:
		err := runner("install", "-v", "github.com/cosmtrek/air@latest")
		if err != nil {
			return err
		}
	case 8:
		err := runner("get", "-v", "github.com/gocolly/colly/v2")
		if err != nil {
			return err
		}
	case 9:
		err := runner("get", "-v", "github.com/microcosm-cc/bluemonday")
		if err != nil {
			return err
		}
	case 10:
		err := runner("get", "-v", "github.com/graphql-go/graphql")
		if err != nil {
			return err
		}
	case 11:
		err := runner("get", "-v", "github.com/arl/statsviz@latest")
		if err != nil {
			return err
		}
	case 12:
		err := runner("get", "-v", "gorm.io/gorm")
		if err != nil {
			return err
		}
	case 13:
		err := runner("get", "-v", "go.uber.org/fx@v1")
		if err != nil {
			return err
		}
	case 14:
		err := runner("get", "-v", "github.com/unidoc/unioffice/")
		if err != nil {
			return err
		}
	case 15:
		err := runner("get", "-v", "github.com/wneessen/go-mail")
		if err != nil {
			return err
		}
	case 16:
		err := runner("install", "-v", "github.com/sqlc-dev/sqlc/cmd/sqlc@latest")
		if err != nil {
			return err
		}
	case 17:
		driver()
	default:
		return errNonePicked

	}
	return nil
}
