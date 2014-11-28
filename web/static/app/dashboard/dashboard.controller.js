(function() {
    'use strict';

    angular
        .module('specter.dashboard')
        .controller('Dashboard', Dashboard);

    Dashboard.$inject = ['$q', 'apiService', 'logger'];

    function Dashboard($q, apiService, logger) {

        /*jshint validthis: true */
        var vm = this;

        vm.news = {
            title: 'Some title...',
            description: 'Some description...'
        };
        vm.title = 'Dashboard';

        activate();

        function activate() {
            var promises = [doNothing()];
//            Using a resolver on all routes or dataservice.ready in every controller
//            return dataservice.ready(promises).then(function(){
            return $q.all(promises).then(function() {
                logger.info('Activated Dashboard View');
            });
        }

        function doNothing() {
            return false;
        }
    }
})();