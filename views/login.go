package views

// THIS IS A GENERATED FILE, EDITS WILL BE OVERWRITTEN
// EDIT THE .haml FILE INSTEAD

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/travissimon/privi/models"
)

func NewLoginWriter() (*LoginWriter) {
	wr := &LoginWriter{}
	
	for idx, pattern := range LoginTemplatePatterns {
		tmpl, err := template.New("LoginTemplates" + string(idx)).Parse(pattern)
		if err != nil {
			fmt.Errorf("Could not parse template: %d", idx)
			panic(err)
		}
		LoginTemplates = append(LoginTemplates, tmpl)
	}
	return wr
}

type LoginWriter struct {
	data *models.LoginResult
}

func (wr *LoginWriter) SetData(data interface{}) {
	wr.data = data.(*models.LoginResult)
}

var LoginHtml = [...]string{
`<html lang="en">
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
		<style type="text/css">
			 body { padding-top: 40px; padding-bottom: 40px; background-color: #f5f5f5; } .form-signin { max-width:
			300px; padding: 19px 29px 29px; margin: 0 auto 20px; background-color: #fff; border: 1px solid #e5e5e5;
			-webkit-border-radius: 5px; -moz-border-radius: 5px; border-radius: 5px; -webkit-box-shadow: 0 1px 2px
			rgba(0,0,0,.05); -moz-box-shadow: 0 1px 2px rgba(0,0,0,.05); box-shadow: 0 1px 2px rgba(0,0,0,.05); }
			.form-signin .form-signin-heading, .form-signin .checkbox { margin-bottom: 10px; } .form-signin input[type="text"],
			.form-signin input[type="password"] { font-size: 16px; height: auto; margin-bottom: 15px; padding: 7px
			9px; }
		</style>
		<div></div>
	</head>
	<body style="zoom: 1;">
		<div class="container">
			<form class="form-signin" method="POST">
				<h2 class="form-signin-heading">Please sign in</h2>
				<input class="input-block-level" type="text" name="txtUsername" placeholder="Username" />
				<input class="input-block-level" type="password" name="txtPassword" placeholder="Password" />
				<label class="checkbox">
					<input type="checkbox" name="chkRememberMe" value="remember-me" /> Remember me
				</label>
				<button class="btn btn-large btn-primary" type="submit">Sign in</button>
				<p><a href="/newuser">Create New Account</a></p>
			</form>
		</div>
		<script src="/static/js/jquery-1.9.1.js"></script>
		<script src="/static/js/bootstrap.min.js"></script>
		
			<script id="hiddenlpsubmitdiv" style="display: none;"></script><script>try{for(var lastpass_iter=0; lastpass_iter
			< document.forms.length; lastpass_iter++){ var lastpass_f = document.forms[lastpass_iter]; if(typeof(lastpass_f.lpsubmitorig2)=="undefined"){
			lastpass_f.lpsubmitorig2 = lastpass_f.submit; lastpass_f.submit = function(){ var form=this; var customEvent
			= document.createEvent("Event"); customEvent.initEvent("lpCustomEvent", true, true); var d = document.getElementById("hiddenlpsubmitdiv");
			for(var i = 0; i < document.forms.length; i++){ if(document.forms[i]==form){ d.innerText=i; } } d.dispatchEvent(customEvent);
			form.lpsubmitorig2(); } } }}catch(e){}</script>
	</body>
</html>
`,
}

var LoginTemplatePatterns = []string{
}

var LoginTemplates = make([]*template.Template, 0, len(LoginTemplatePatterns))

func (wr LoginWriter) Execute(w http.ResponseWriter, r *http.Request) {
	wr.ExecuteData(w, r, wr.data)
}

func (wr *LoginWriter) ExecuteData(w http.ResponseWriter, r *http.Request, data *models.LoginResult) {
	var err error = nil
	fmt.Fprint(w, LoginHtml[0])
if err != nil {err = nil}}

func handleLoginError(err error) {
	if err != nil {fmt.Println(err)}}