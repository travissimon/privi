package views

// THIS IS A GENERATED FILE, EDITS WILL BE OVERWRITTEN
// EDIT THE .haml FILE INSTEAD

import (
	"fmt"
	"html/template"
	"net/http"
)

func NewHomeWriter() (*HomeWriter) {
	wr := &HomeWriter{}
	
	for idx, pattern := range HomeTemplatePatterns {
		tmpl, err := template.New("HomeTemplates" + string(idx)).Parse(pattern)
		if err != nil {
			fmt.Errorf("Could not parse template: %d", idx)
			panic(err)
		}
		HomeTemplates = append(HomeTemplates, tmpl)
	}
	return wr
}

type HomeWriter struct {
	data interface{}
}

func (wr *HomeWriter) SetData(data interface{}) {
	wr.data = data.(interface{})
}

var HomeHtml = [...]string{
`<doctype />
<html lang="en">
	<head>
		<title>Privi Login</title>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8"></meta>
		<meta charset="utf-8"></meta>
		<meta name="viewport" content="width=device-width" initial-scale="1.0"></meta>
		<link href="/static/css/bootstrap.min.css" rel="stylesheet"></link>
		<link href="/static/css/bootstrap-responsive.css"></link>
		<!-- HTML5 shim, for IE6-8 support of HTML5 elements -->
		<!--[if lt IE 9]> <script src="../assets/js/html5shiv.js"></script>
		<![endif]-->
		<style type="text/css"></style>
		html { height:100%;  max-height:100%; 
		
			} body { height:100%;  max-height:100%;  padding-top: 60px; padding-bottom: 40px; } .sidebar-nav { padding:
			9px 0; } @media (max-width: 980px) { .navbar-text.pull-right { float: none; padding-left: 5px; padding-right:
			5px; } }
			<div></div>
	</head>
	<body style="zoom: 1;">
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
						<p class="navbar-text pull-right"> Logged in as <a href="/userdetails/tsimon" class="navbar-link">tsimon</a></p>
						<ul class="nav">
							<li><a href="/mail">Mail</a></li>
							<li><a href="/news">News</a></li>
							<li><a href="/community">Community</a></li>
						</ul>
					</div>
				</div>
			</div>
		</div>
		<!-- Home markup -->
		<div class="container">
			<div class="hero-unit">
				<h1>Welcome to Privi</h1>
				<p>
					 Welcome to the Privi website! Privi is intended to be an NSA-free replacement for many of your online
					activities, including email, rss and social networking.
				</p>
			</div>
			<div class="row">
				<div class="span4">
					<h2>Email</h2>
					<p>
						 Access your email securely from anywhere in the world using our web-based interface. All accesses to
						your email goes through SSL, keeping it safe-and- secure.
					</p>
				</div>
				<div class="span4">
					<h2>News</h2>
					<p>
						 Without Google shutting down Reader, we never would have built this site. Use our News feature to browse
						RSS feeds and share items with other users, just like the old Reader used to do.
					</p>
				</div>
				<div class="span4">
					<h2>Social Networking</h2>
					<p>
						 Are you concerned about Facebook providing all of your information to various Governments? Or are you
						just sick of Facebook? We are too! Use this instead. It may be crap, but it's out crap.
					</p>
				</div>
			</div>
			<hr></hr>
			<footer>
				<p>© Privi 2013</p>
			</footer>
		</div>
		<script src="/static/js/jquery-1.9.1.js"></script>
		<script src="/static/js/bootstrap.min.js"></script>
	</body>
</html>
`,
}

var HomeTemplatePatterns = []string{
}

var HomeTemplates = make([]*template.Template, 0, len(HomeTemplatePatterns))

func (wr HomeWriter) Execute(w http.ResponseWriter, r *http.Request) {
	wr.ExecuteData(w, r, wr.data)
}

func (wr *HomeWriter) ExecuteData(w http.ResponseWriter, r *http.Request, data interface{}) {
	var err error = nil
	fmt.Fprint(w, HomeHtml[0])
if err != nil {err = nil}}

func handleHomeError(err error) {
	if err != nil {fmt.Println(err)}}