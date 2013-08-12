angular.module('news', ['ui.bootstrap']);

function NewsController($scope, $http) {
  $scope.feedName = 'News';
  $scope.entries = [];

  $scope.loadFeed = function(feedId, feedName) {
	  $http.get('/news/' + feedId).success(function(data) {
		  $scope.entries = data;
		  $scope.feedName = feedName;
	  });
  }

}
