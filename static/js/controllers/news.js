angular.module('news', ['ui.bootstrap']);

function NewsController($scope, $http) {
  $scope.feedName = 'News';
  $scope.entries = [];
  $scope.newFeedURL = "";

  $scope.loadFeed = function(feedId, feedName) {
	  $http.get('/news/' + feedId).success(function(data) {
		  $scope.entries = data;
		  $scope.feedName = feedName;
	  });
  }

  $scope.error = function(errorMsg) {
	  alert(errorMsg);
  }

  $scope.addFeed = function() {
	  var url = $scope.newFeedURL;
	  if (url == null || url.length == 0) {
		  $scope.error("URL is empty");
		  return
	  }
	  var data = {feedUrl: url};
	  $http.post('/news/newfeed', data).success(function(data) {
		  if (data.ErrorMessage == nil || data.ErrorMessage.length == 0) {
			  $scope.entries = data.Entries;
			  $scope.feedName = data.FeedName
		  } else {
		  }
		  $scope.error("returned from server: " + data.ErrorMessage);
	  });
  }

}
