news = angular.module('news', ['ui.bootstrap']);

function NewsController($scope, $http) {
  $scope.pendingFeedName = "";
  $scope.feedName = 'News';
  $scope.feeds = [];
  $scope.entries = [];
  $scope.newFeedURL = "";

  $scope.loadFeeds = function() {
	  $http.get('/news/allFeeds').success(function(data) {
		  $scope.feeds = data;
	  });
  }

  $scope.loadFeed = function(feedId, feedName) {
	  $scope.pendingFeedName = feedName
	  $http.get('/news/' + feedId).success(function(data) {
		  $scope.entries = data;
		  $scope.feedName = $scope.pendingFeedName;
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
		  $scope.error("returned from server: " + JSON.stringify(data));
		  if (data.ErrorMessage == nil || data.ErrorMessage.length == 0) {
			  var newFeed = {
				  Id: data.FeedId,
				  Title: data.FeedName,
			  };
			  // find alphebetised insertion point
			  var idx = 0;
			  while (idx < $scope.feeds.length && newFeed.Title < $scope.feeds[i].Title)
				  idx++;

			  $scope.feeds.splice(idx, 0, newFeed);
		  } else {
		  }
		  $scope.error("returned from server: " + data.ErrorMessage);
	  });
  }

}
