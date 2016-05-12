This is a mini search engine based on ElasticSearch(https://www.elastic.co/).

There are two main functions in this program.
1) crawling web pages from a starting url.
Crawling is executed in cigol-crawler.go. A starting url can be passed to the program or a default one is set in the code. The program will crawl the pages and store links that are in the crawled pages for next crawling. The crawled pages will be stored into the configed ES servers.

2) Search engine service.
cigol-search-engine.go can start a http-server to supply services for searching pages crawled in step 1). Users can search against web page title or content. Other advanced queries will come soon...