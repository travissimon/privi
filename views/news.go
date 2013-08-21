package views

// THIS IS A GENERATED FILE, EDITS WILL BE OVERWRITTEN
// EDIT THE .haml FILE INSTEAD

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/travissimon/privi/models"
)

func NewNewsWriter() (*NewsWriter) {
	wr := &NewsWriter{}
	
	for idx, pattern := range NewsTemplatePatterns {
		tmpl, err := template.New("NewsTemplates" + string(idx)).Parse(pattern)
		if err != nil {
			fmt.Errorf("Could not parse template: %d", idx)
			panic(err)
		}
		NewsTemplates = append(NewsTemplates, tmpl)
	}
	return wr
}

type NewsWriter struct {
	data *models.NewsPage
}

func (wr *NewsWriter) SetData(data interface{}) {
	wr.data = data.(*models.NewsPage)
}

var NewsHtml = [...]string{
`<doctype />
<html lang="en" ng-app="news">
	<head>
		<title>Privi News</title>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8"></meta>
		<meta charset="utf-8"></meta>
		<meta name="viewport" content="width=device-width" initial-scale="1.0"></meta>
		<link href="/static/css/bootstrap.min.css" rel="stylesheet"></link>
		<link href="/static/css/bootstrap-responsive.css" rel="stylesheet"></link>
		<style type="text/css">
			 html { height:100%;  max-height:100%;  } body { height:100%;  max-height:100%;  padding-top: 60px; padding-bottom:
			40px; } .sidebar-nav { padding: 9px 0; } @media (max-width: 980px) { .navbar-text.pull-right { float:
			none; padding-left: 5px; padding-right: 5px; } }
		</style>
	</head>
	<body style="zoom: 1;" ng-controller="NewsController">
		<div class="navbar navbar-inverse navbar-fixed-top">
			<div class="navbar-inner">
				<div class="container-fluid">
					<button class="btn btn-navbar" type="button" data-toggle="collapse" data-target=".nav-collapse">
						<span class="icon-bar"></span>
						<span class="icon-bar"></span>
						<span class="icon-bar"></span>
					</button>
					<a class="brand" href="/">Privi</a>
					<div class="nav-collapse collapse">
						<p class="navbar-text pull-right">
							`,
							`
						</p>
						<ul class="nav">
							<li><a href="/mail">Mail</a></li>
							<li class="active"><a href="/news">News</a></li>
							<li><a href="/community">Community</a></li>
						</ul>
					</div>
				</div>
			</div>
		</div>
		<!-- Home markup -->
		<div class="container-fluid">
			<div class="row-fluid">
				<div class="span3">
					<div class="well sidebar nav">
						<ul class="nav nav-list">
							<li class="nav-header">Feeds</li>
							`,
							`
							<li>
								`,
								`
							</li>
							<li class="nav-header">Options</li>
							<li>
								<input type="text" placeholder="Feed URL" ng-model="newFeedURL"></input>
								<br />
								<button class="btn btn-small" ng-click="addFeed()"><i class="icon-plus"></i> Add Feed</button>
							</li>
						</ul>
					</div>
				</div>
				<div class="span9">
					<div class="row-fluid">
						<h3>{{ feedName }}</h3>
						<accordion close-others="true">
							<accordion-group heading="{{entry.Title}}" ng-repeat="entry in entries">
								<div ng-bind-html-unsafe="entry.Summary"></div>
							</accordion-group>
						</accordion>
					</div>
				</div>
			</div>
			<hr></hr>
			<footer>
				<p>© Privi 2013</p>
			</footer>
		</div>
		<script src="/static/js/angular.min.js"></script>
		<script src="/static/js/ui-bootstrap-tpls-0.5.0.min.js"></script>
		<script src="/static/js/controllers/news.js"></script>
	</body>
</html>
`,
}

var NewsTemplatePatterns = []string{
	`Logged in as <a href='/userdetails/{{.Username }}' class='navbar-link'>{{.Username}}</a>`,
	`<a ng-click='loadFeed({{.Id}})'>{{.Title}}</a>`,
}

var NewsTemplates = make([]*template.Template, 0, len(NewsTemplatePatterns))

func (wr NewsWriter) Execute(w http.ResponseWriter, r *http.Request) {
	wr.ExecuteData(w, r, wr.data)
}

func (wr *NewsWriter) ExecuteData(w http.ResponseWriter, r *http.Request, data *models.NewsPage) {
	var err error = nil
	fmt.Fprint(w, NewsHtml[0])
	err = NewsTemplates[0].Execute(w, data)
	handleNewsError(err)
	fmt.Fprint(w, NewsHtml[1])
	for _, feed := range data.Feeds {
		fmt.Fprint(w, NewsHtml[2])
		err = NewsTemplates[1].Execute(w, feed)
		handleNewsError(err)
	}
	fmt.Fprint(w, NewsHtml[3])
if err != nil {err = nil}}

func handleNewsError(err error) {
	if err != nil {fmt.Println(err)}}