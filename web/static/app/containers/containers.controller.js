(function() {
    'use strict';

    angular
        .module('specter.containers')
        .controller('Containers', Containers);

    /* @ngInject */
    function Containers(dataservice, logger) {
        /*jshint validthis: true */
        var vm = this;
        vm.containers = [];
        vm.title = 'Containers';

        activate();

        function activate() {
//            Using a resolver on all routes or dataservice.ready in every controller
//            var promises = [getAvengers()];
//            return dataservice.ready(promises).then(function(){
            return getContainers().then(function() {
                logger.info('Activated Containers View');
            });
        }

        function getContainers() {
            return dataservice.getContainers().then(function(data) {
                vm.containers = data;
                return vm.containers;
            });
        }
    }
})();