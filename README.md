# geocoder

Who's On First-focused coarse multi-lingual geocoder.

## Motivation

These tools are modeled after the [pelias/placeholder](https://github.com/pelias/placeholder) project. SFO Museum needed something written in Go, with support for the exiting Go language Who's On First tooling and the ability to apply custom filtering (see below). So now this package exists.

## Almost stable

This package and the tools it provides are "almost stable". Which is to say there probably won't be any major changes but there _might_ be. Mostly the question centers on whether or not extra columns needed to be added to certain database tables for improved sorting.

This accounts for the initial `v0.9.0` version number. Once things have settled down it will be promoted to `v1.0.0`.

## Geocoding support

Currently only "coarse" (or admin-level) geocoding is supported. This does not include venues yet but will, eventually. I am not thinking about address-level geocoding at this time.

### Query filters

Query strings may be filtered by the following:

* One or more 2-letter ISO country codes.
* One or more Who's On First placetypes or custom `wof:placetype_alt` placetypes.
* Start and end dates represented by ExtendedDateTimeFormat (EDTF) date strings.
* Custom bounding box.
* One or more Who's On First IDs which are ancestors (contain) result candidates.
* 3-letter language code for a place name.
* Who's On First language "tag" used to signal Who's On First classify a place name (for example: preferred, variant, historical, etc.)

## Database support

Currently only SQLite databases are supported. The design of the code (interfaces) is such that it should be possible to add other databases but that work has not happened yet.

Database tables are created at runtime, if necessary, so you don't need to do that manually. Database schemas can be found in [coarse/sql/*.schema](coarse/sql). Note that for reasons of efficiency none of the individual table schemas have indices. Additionally, by default, any existing indices and FTS-related tables are _removed_ before indexing is begun and then (re)created once indexing is complete. This behaviour can be modified but you'll need to be explicit about doing so.

Finally, the "post-indexing" phase currently takes longer than I would like to account for the need to de-duplicate rows in the `ancestors` and `placetypes_alt` tables before a unique indices are created. _This should not be necessary._ It suggests that the same Who's On First record is contained by two or more data sources (repositories) which is an error. That will be corrected in time but, until then, this extra step (and time) is necessary.

## Database URIs

Geocoding databases are specified using the `-geocoder-uri` flag which define database specifics in the form of a URI.

### SQLite

SQLite geocoder databases take the form of:

```
sql://sqlite?dsn={SQLITE_DSN_STRING}
```

For example:

```
sql://sqlite?dsn=wof.db
```

## Tools

```
$> make cli
go build -mod vendor -ldflags="" -o bin/wof-coarse-geocoder-index cmd/wof-coarse-geocoder-index/main.go
go build -mod vendor -ldflags="" -o bin/wof-coarse-geocoder-query cmd/wof-coarse-geocoder-query/main.go
go build -mod vendor -ldflags="" -o bin/wof-coarse-geocoder-server cmd/wof-coarse-geocoder-server/main.go
```

### wof-coarse-geocoder-index

Index one or more Who's On First data sources in a (coarse) geocoding database.

```
$> ./bin/wof-coarse-geocoder-index -h
Index one or more Who's On First data sources in a (coarse) geocoding database.
Usage:
	./bin/wof-coarse-geocoder-index [options] uri(N) uri(N) uri(N)
Valid options are:
  -exclude-deprecated
    	Do not index records which have been deprecated. (default true)
  -exclude-funky
    	Do not index records which have been flagged as "funky". (default true)
  -exclude-nullisland
    	Do not index records that are "visiting" Null Island (have 0,0 coordinate data). (default true)
  -exclude-superseded
    	Do not index records which have been superseded. (default true)
  -fresh
    	This flags signals that a fresh database is being indexed disabling checks for existing or updated records.
  -geocoder-uri string
    	A registered whosonfirst/geocoder/coarse.Geocoder URI. (default "sql://sqlite?dsn=:memory:")
  -index-juggling
    	Perform indexing speed optiomizations. This will include dropping existing indices and the FTS table prior to indexing and (re)adding them at the end. (default true)
  -iterator-uri string
    	A registered whosonfirst/go-whosonfirst/v4/iterate.Iterate URI. (default "repo://")
  -offset int
    	Optional document offset to start indexing from.
  -prune
    	Prune existing records before (re)adding them to the database.
  -verbose
    	Enable verbose (debug) logging.
```

For example:

```
$> ./bin/wof-coarse-geocoder-index \
	-fresh \
	-iterator-uri repo:// \
	-geocoder-uri 'sql://sqlite?dsn=sfom.db' \
	/usr/local/data/sfomuseum-data-architecture \
	/usr/local/data/sfomuseum-data-whosonfirst
	
2026/08/08 11:15:04 INFO Rewrote iterator URI uri="repo:?exclude=properites.edtf%3Adeprecated%3D.%2A&exclude=properites.wof%3Asuperseded_by%3D.%2A&exclude=properites.mz%3Ais_funky%3D1"
2026/08/08 11:15:04 INFO Pre-indexing complete time=154.5µs
2026/08/08 11:16:04 INFO Iterator stats elapsed=1m0.001052708s seen=2530 allocated="9.7 MB" "total allocated"="498 MB" sys="54 MB" numgc=103
2026/08/08 11:16:04 INFO Indexing stats elapsed=1m0.001189958s seen=2530 "average (ms)"=0.09881422924901186

...time passes

2026/08/08 11:21:04 INFO Iterator stats elapsed=6m0.00110525s seen=3469 allocated="37 MB" "total allocated"="4.2 GB" sys="275 MB" numgc=287
2026/08/08 11:21:33 INFO Iterator stats elapsed=6m29.009900916s seen=3623 allocated="60 MB" "total allocated"="4.4 GB" sys="275 MB" numgc=294
2026/08/08 11:21:33 INFO Indexing complete seen=3623 time=6m29.014973666s "average (ms)"=0.2555892906431134
2026/08/08 11:21:34 INFO Post-indexing complete time=661.16725ms "time (total)"=6m29.676154166s
```

Indexing time can depend a lot on the data source. Files on disk (above) can take a while. Indexing Who's On First Parquet files (produced by the [wof-parquet-export](https://github.com/whosonfirst/go-whosonfirst) tool in the `whosonfirst/go-whosonfirst` package) is significantly faster, taking only 20-30 minutes to create a geocoding database for all 6 million plus records:

```
$> ./bin/wof-coarse-geocoder-index \
	-fresh \
	-iterator-uri parquet:// \
	-geocoder-uri 'sql://sqlite?dsn=wof-sfom.db' \
	/usr/local/data/whosonfirst-parquet/whosonfirst-data-admin-*.parquet
	
2026/08/08 11:31:00 INFO Rewrote iterator URI uri="parquet:?exclude=properites.edtf%3Adeprecated%3D.%2A&exclude=properites.wof%3Asuperseded_by%3D.%2A&exclude=properites.mz%3Ais_funky%3D1"
2026/08/08 11:31:00 INFO Pre-indexing complete time=88.792µs
2026/08/08 11:32:00 INFO Iterator stats elapsed=1m0.000199625s seen=319775 allocated="3.2 GB" "total allocated"="38 GB" sys="4.5 GB" numgc=88
2026/08/08 11:32:00 INFO Indexing stats elapsed=1m0.000307583s seen=319765 "average (ms)"=0.09237721451691086

...time passes

2026/08/08 11:51:01 INFO Iterator stats elapsed=20m1.057480125s seen=6457706 allocated="4.1 GB" "total allocated"="576 GB" sys="19 GB" numgc=374

...more time passes

2026/08/08 12:01:46 INFO Post-indexing complete time=10m44.596903375s "time (total)"=30m45.71374225s

```

That database, in turn, can be supplemented with SFO Museum specific Who's On First style data repositories. For example:

```
$> ./bin/wof-coarse-geocoder-index \
	-prune \
	-iterator-uri repo:// \
	-geocoder-uri 'sql://sqlite?dsn=wof-sfom.db' \
	/usr/local/data/sfomuseum-data-architecture \
	/usr/local/data/sfomuseum-data-whosonfirst
	
2026/08/08 12:33:40 INFO Rewrote iterator URI uri="repo:?exclude=properites.edtf%3Adeprecated%3D.%2A&exclude=properites.wof%3Asuperseded_by%3D.%2A&exclude=properites.mz%3Ais_funky%3D1"
2026/08/08 12:33:40 INFO Pre-indexing complete time=95.542µs
2026/08/08 12:34:40 INFO Iterator stats elapsed=1m0.001063042s seen=1229 allocated="5.2 MB" "total allocated"="262 MB" sys="46 MB" numgc=60
2026/08/08 12:34:40 INFO Indexing stats elapsed=1m0.001180792s seen=1229 "average (ms)"=0.10903173311635476

...time passes

2026/08/08 12:42:25 INFO Iterator stats elapsed=8m44.648322417s seen=3623 allocated="43 MB" "total allocated"="4.4 GB" sys="317 MB" numgc=295
2026/08/08 12:42:25 INFO Indexing complete seen=3623 time=8m44.654067s "average (ms)"=0.281810654154016
...
2026/08/08 12:48:44 INFO Post-indexing complete time=6m19.407913166s "time (total)"=15m4.061988625s

$> du -h wof-sfom.db 
6.5G	wof-sfom.db
```

Valid data sources are anything the [whosonfirst/go-whosonfirst/v4/iterate](https://github.com/whosonfirst/go-whosonfirst/tree/main/iterate) package can support. Please consult documentation for details.

### wof-coarse-geocoder-query

Query a Who's On First (coarse) geocoding database.

```
$> ./bin/wof-coarse-geocoder-query -h
Query a Who's On First (coarse) geocoding database.
Usage:
	./bin/wof-coarse-geocoder-query [options]
Valid options are:
  -belongs-to value
    	Zero or more Who's On First ancestor IDs to filter results by.
  -bounds string
    	Optional bounding box (in the form of 'minx,miny,maxx,mayx') to filter results by.
  -country value
    	Zero or more 2-letter country codes to filter results by.
  -date-ends string
    	Optional ETDF ending date string to filter results by.
  -date-starts string
    	Optional ETDF starting date string to filter results by.
  -geocoder-uri string
    	A registered whosonfirst/geocoder/coarse.Geocoder URI.
  -lang string
    	An optional (3-letter) language code to filter results by,
  -mode string
    	Output mode for results. Valid options are: geojson, tab. (default "tab")
  -page int
    	The specific page number to query for paginated result sets. (default 1)
  -per-page int
    	The number of results to include for paginated result sets. (default 100)
  -placetype value
    	Zero or more placetypes to filter results by.
  -query string
    	The term to query for. Required.
  -tag string
    	An option WOF language tag to filter results by.
  -verbose
    	Enable verbose (debug) logging.
```

For example:

```
$> ./bin/wof-coarse-geocoder-query \
	-geocoder-uri 'sql://sqlite?dsn=wof.db' \
	-query T3

2026/08/08 11:25:22 INFO Query results total=7 page=1 pages=1

id		name		placetype	is current	inception	cessation	label
1947304447	Terminal 3	wing		1		2024-11-05	..		Terminal 3, SFO Terminal Complex, San Francisco International Airport, San Francisco, US
1159157307	Terminal 3	wing		0		2017~		2019-07-23	Terminal 3, SFO Terminal Complex, San Francisco International Airport, San Francisco, US
1477855699	Terminal 3	wing		0		2019-07-23	2020-~05	Terminal 3, SFO Terminal Complex, San Francisco International Airport, San Francisco, US
1729792487	Terminal 3	wing		0		2020-~05	2021-05-25	Terminal 3, SFO Terminal Complex, San Francisco International Airport, San Francisco, US
1745882233	Terminal 3	wing		0		2021-05-25	2021-11-09	Terminal 3, SFO Terminal Complex, San Francisco International Airport, San Francisco, US
1763588269	Terminal 3	wing		0		2021-11-09	2024-06-17	Terminal 3, SFO Terminal Complex, San Francisco International Airport, San Francisco, US
1914600841	Terminal 3	wing		0		2024-06-17	2024-11-05	Terminal 3, SFO Terminal Complex, San Francisco International Airport, San Francisco, US
```

Or to query with a custom placetype (stored in the `wof:placetype_alt` property):

```
$> ./bin/wof-coarse-geocoder-query \
	-geocoder-uri 'sql://sqlite?dsn=wof-sfom.db' \
	-query SFO \
	-placetype airport

2026/08/08 11:27:37 INFO Query results total=1 page=1 pages=1

id		name					placetype	is current	inception	cessation	label
102527513	San Francisco International Airport	campus		1		1948~		..		San Francisco International Airport, San Francisco, California, US
```

You can also query for records using a known concordances, for example an IATA airport code:

```
$> ./bin/wof-coarse-geocoder-query \
	-geocoder-uri 'sql://sqlite?dsn=wof-sfom.db' \
	-query iata:code=YUL
	
2026/08/08 13:06:39 INFO Query results total=1 page=1 pages=1

id		name							placetype	is current	inception	cessation	label
102554351	Montreal-Pierre Elliott Trudeau International Airport	campus		1		1941-09-01			Montreal-Pierre Elliott Trudeau International Airport, Dorval, Quebec, CA
```

Or a GeoPlanet identifier:

```
$> ./bin/wof-coarse-geocoder-query \
	-geocoder-uri 'sql://sqlite?dsn=wof-sfom.db' \
	-query gp:id=27978
	
2026/08/08 13:09:26 INFO Query results total=2 page=1 pages=1

id		name		placetype	is current	inception	cessation	label
101750367	London		locality	1		0043~				London, Greater London, GB
1880762729	Greater London	region		1						Greater London, GB
```

The geocoder will pass the so-called "Brooklyn test" in English:

```
$> ./bin/wof-coarse-geocoder-query \
	-geocoder-uri 'sql://sqlite?dsn=wof-sfom.db' \
	-query brooklyn \
	-per-page 10
	
2026/08/09 22:20:13 INFO Query results total=135 page=1 pages=14

id		name			placetype	is current	inception	cessation	label
421205765	Brooklyn		borough		1						Brooklyn, New York, New York, US
85969229	Brooklyn Park		locality	1						Brooklyn Park, Minnesota, US
404511829	Brooklyn Park		localadmin	1						Brooklyn Park, Minnesota, US
85807925	Brooklyn Heights	neighbourhood	1						Brooklyn Heights, New York, New York, US
85871819	Old Brooklyn		neighbourhood	1						Old Brooklyn, Cleveland, Cleveland, Ohio, US
85969235	Brooklyn Center		locality	1						Brooklyn Center, Minnesota, US
404511827	Brooklyn Center		localadmin	1						Brooklyn Center, Minnesota, US
101712549	Brooklyn		locality	1						Brooklyn, Ohio, US
85949701	Brooklyn Park		locality	1						Brooklyn Park, Maryland, US
404525053	Brooklyn		localadmin	1						Brooklyn, Ohio, US
```

And in other languages, like Farsi:

```
$> ./bin/wof-coarse-geocoder-query \
	-geocoder-uri 'sql://sqlite?dsn=wof-sfom.db' \
	-query بروکلین \
	-per-page 10
	
2026/08/09 22:22:22 INFO Query results total=72 page=1 pages=8

id		name			placetype	is current	inception	cessation	label
421205765	Brooklyn		borough		1						Brooklyn, New York, New York, US
85969229	Brooklyn Park		locality	1						Brooklyn Park, Minnesota, US
85807925	Brooklyn Heights	neighbourhood	1						Brooklyn Heights, New York, New York, US
85871819	Old Brooklyn		neighbourhood	1						Old Brooklyn, Cleveland, Cleveland, Ohio, US
85969235	Brooklyn Center		locality	1						Brooklyn Center, Minnesota, US
101712549	Brooklyn		locality	1						Brooklyn, Ohio, US
85949701	Brooklyn Park		locality	1						Brooklyn Park, Maryland, US
404525053	Brooklyn		localadmin	1						Brooklyn, Ohio, US
404495913	Brooklyn		localadmin	1						Brooklyn, Connecticut, US
85807887	Brooklyn		neighbourhood	1						Brooklyn, Jacksonville, Florida, US
```

### wof-coarse-geocoder-server

HTTP server for handling requests against a Who's On First (coarse) geocoding database.

```
$> ./bin/wof-coarse-geocoder-server -h
HTTP server for handling requests against a Who's On First (coarse) geocoding database.
Usage:
	./bin/wof-coarse-geocoder-server [options]
Valid options are:
  -demo
    	Start a web-based demo on the root URL of the server.
  -geocoder-uri string
    	A registered whosonfirst/geocoder/coarse.Geocoder URI.
  -pagination-per-page int
    	The maximum number of results to include per API request. (default 50)
  -prefix string
    	An optional URL prefix to listen for requests on.
  -server-uri string
    	A registered aaronland/go-http/v4/server.Server URI. (default "http://localhost:8080")
  -verbose
    	Enable verbose (debug) logging.
```	

For example:

```
$> make server GEOCODER_URI='sql://sqlite?dsn=test.db'
go run -mod vendor cmd/wof-coarse-geocoder-server/main.go \
		-verbose \
		-server-uri http://localhost:8080 \
		-geocoder-uri sql://sqlite?dsn=test.db
2026/08/05 10:05:35 DEBUG Verbose logging enabled
2026/08/05 10:05:35 INFO Listening for requests address=http://localhost:8080

2026/08/05 10:07:23 DEBUG Time to query query=SFO time=3.302416ms
```

And then:

```
$> curl -s 'http://localhost:8080/api/query/?query=SFO&placetype=airport' | jq
{
  "pagination": {
    "total": 1,
    "per_page": 50,
    "page": 1,
    "pages": 1,
    "next_page": 0,
    "previous_page": 0
  },
  "results": {
    "features": [
      {
        "id": 102527513,
        "type": "Feature",
        "bbox": [
          -122.408061,
          37.601617,
          -122.354907,
          37.640167
        ],
        "geometry": {
          "type": "Point",
          "coordinates": [
            -122.370943,
            37.61799
          ]
        },
        "properties": {
          "edtf:cessation": "..",
          "edtf:inception": "1948~",
          "geocoder:rank": -13.013866678659102,
          "mz:is_current": 1,
          "wof:country": "US",
          "wof:hierarchies": [
            {
              "campus_id": 102527513,
              "continent_id": 102191575,
              "country_id": 85633793,
              "county_id": 102087579,
              "locality_id": 85922583,
              "postalcode_id": 554784711,
              "region_id": 85688637
            },
            {
              "campus_id": 102527513,
              "continent_id": 102191575,
              "country_id": 85633793,
              "county_id": 102085387,
              "region_id": 85688637
            }
          ],
          "wof:id": 102527513,
          "wof:label": "San Francisco International Airport, San Francisco, California, US",
          "wof:name": "San Francisco International Airport",
          "wof:parent_id": 85922583,
          "wof:placetype": "campus",
          "wof:placetype_alt": [
            "airport"
          ]
        }
      }
    ],
    "type": "FeatureCollection"
  }
}
```

#### Demo mode

When started with the `-demo` flag the server will host a simple web application at its root URL. When you open your web browser to `http://localhost:8080` (or whatever you've configured the `-server-uri` flag to be) you'll see something like this:

![](docs/images/geocoder-demo-launch.png)

By default you can enter a query term in the search box in the top-right hand corner but you can also perform a more detailed query by opening the "Advanced" menu beneath the map:

![](docs/images/geocoder-demo-advanced.png)

Query results are displayed on the map and in a list view beneath the map:

![](docs/images/geocoder-demo-paris.png)

Clicking on either a location's point on the map or its ID in list view will display its map label and jump to the record in the list view.

![](docs/images/geocoder-demo-paris-2.png)

That's all it does for the time being.

## Data sources

Currently, only Who's On First-shaped documents are supported. Internally those documents are transformed in an internal `Record` struct which looks like this:

```
type Record struct {
	Id             int64                          `json:"wof:id"`
	ParentId       int64                          `json:"wof:parent_id"`
	Name           string                         `json:"wof:name"`
	Country        string                         `json:"wof:country"`
	Placetype      string                         `json:"wof:placetype"`
	PlacetypeAlt   []string                       `json:"wof:placetype_alt"`
	Hierarchies    []map[string]int64             `json:"wof:hierarchies"`
	Centroid       *orb.Point                     `json:"wof:centroid"`
	Bounds         []orb.Bound                    `json:"wof:bounds"`
	Inception      string                         `json:"edtf:inception,omitempty"`
	Cessation      string                         `json:"etdf:cessation,omitempty"`
	PopulationRank int64                          `json:"wof:population_rank,omitempty"`
	IsCurrent      string                         `json:"mz:is_current,omitempty"`
	Tokens         map[string]map[string][]string `json:"tokens,omitempty"`
}
```

Going forward the "easiest" thing may be to simply change this data structure to assume that all identifiers are strings and do the extra work, internally, to convert them to and from integer values. Maybe? It's just too soon to think about right now.