# stocksbeat

Welcome to stocksbeat.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/tuxiedev/stocksbeat`


## Quickstart
To run this application, you will need an API key from [Finnhub](https://finnhub.io). Please follow their documentation on how to get an API. To summarise, you will need to do the following steps
* [Register](https://finnhub.io/register) for a free account
* [Log in](https://finnhub.io/login) to your account dashboard.
* You will find your account key on the first page after you login. Take a note of that API key, you will need it later.

Head over to the [Releases](https://github.com/tuxiedev/stocksbeat/releases) and download the version required appropriate for your system. 


### Configuration
No matter what version you download, the configuration for Beat will be done by `stocksbeat.yml` file. This file can exist in either your `/etc/stocksbeat` directory or will be available in the unzipped/untared location. All you have to do now is to list the stocks you want to consume data for and configure where you want the data be written to.

You can either choose to edit the `stocksbeat.yml` file or create your own YAML and follow on

#### Configure Elasticsearch and Kibana
Elastic Beats support a wide range of outputs. In this example, I will only configure the beats to write to a running Elasticsearch cluster.

If you have Docker and Docker-Compose on your machine, you can use the handly docker-compose.yml that comes with this repository. You can either choose to clone the repo, or just download the [docker-compose.yml](https://raw.githubusercontent.com/tuxiedev/stocksbeat/master/docker-compose.yml) to your local and start the stack using:

```shell
$ docker-compose up
```
After Elasticsearch and Kibana have both started on your system, you should be able to see the Kibana UI on http://localhost:5601.

Now, lets add the configuration to our output in the `stocksbeat.yml` file.
```yaml
# stocksbeat.yml
output:
    elasticsearch:
        hosts: ["localhost:9200"]
```
Note that there are plethora of options to send the trades to. I am just going to leave a link to the [Elastic Filebeat documentation](https://www.elastic.co/guide/en/beats/filebeat/current/configuring-output.html). Just whereever that documentation says `filebeat.yml`, in this context, it will really mean `stocksbeat.yml`. 

Alright, so now that we have a place to store the trades we consume, we are ready to pick some stocks to follow trades of. For the sake of this example, I am going to choose all stocks in the [FAANG](https://www.investopedia.com/terms/f/faang-stocks.asp) group. To follow those trades, configure this stanza in `stocksbeat.yml`.
```yaml
# stocksbeat.yml
stocksbeat:
  symbols:
    - FB   # F-acebook
    - AMZN # A-mazon
    - AAPL # A-apple
    - NFLX # N-etflix
    - GOOG # A-lphabet (wait, what? well, it was formerly G-oogle)
  finnhubToken: ${FINNHUB_TOKEN}
```
Notice why we put the value for `finnhubToken` in a special format? Well, that the method you'd use to set the value of that configuration to something coming from envrionment variables. We do that so we don't accidently commit the token or keep it written hardcoded in plaintext on a file.

Alright now that our configuration file is ready, it should look like the following
```yaml
stocksbeat:
    symbols:
        - FB   # F-acebook
        - AMZN # A-mazon
        - AAPL # A-apple
        - NFLX # N-etflix
        - GOOG # A-lphabet (wait, what? well, it was formerly G-oogle)
    finnhubToken: ${FINNHUB_TOKEN}
output:
    elasticsearch:
        hosts: ["localhost:9200"]
```
Let's keep the file aside for now. We have to warmup the Elasticsearch cluster with [Index Templates](https://www.elastic.co/guide/en/elasticsearch/reference/current/index-templates.html)
> An index template is a way to tell Elasticsearch how to configure an index when it is created.

We generally definte index templates so that Elasticsearch fully knows what type of data to expect for certain fields in the document. Index templates are also used to set default settings for index at creation time, like the number of shards, number of replicas and [lot more](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-update-settings.html).

So, lets install that template
```shell
./stocksbeat -e setup -c stocksbeat.yml
```
This command will install the index template on the output Elasticsearch cluster we configured earlier.

> Come on! Let me see some data already!

Alright alright! Hold on.. one more thing

> what?

...

Nothing let's start

```
./stocksbeat -e
```
Notice why we are using the `-e` flag all the time? Well, its just so that logs are printed on the console too. You can run all the commands without `-e` flag, and the logs will be stored by default in the `logs` directory.


As data starts indexing, we can head over to some Kibana dashboards to look at what all we are collecting



## Development

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Clone

To clone stocksbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/tuxiedev/stocksbeat
git clone https://github.com/tuxiedev/stocksbeat ${GOPATH}/src/github.com/tuxiedev/stocksbeat
```

### Init Project
To get running with stocksbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

### Build

To build the binary for stocksbeat run the command below. This will generate a binary
in the same directory with the name stocksbeat.

```
make
```


### Run

To run stocksbeat with debugging output enabled, run:

```
FINNHUB_TOKEN=<your_finnhub_token> ./stocksbeat -c stocksbeat.yml -e -d "*"
```


### Cleanup

To clean  stocksbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
