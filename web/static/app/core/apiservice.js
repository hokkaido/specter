(function() {
    'use strict';

    angular
        .module('specter.core')
        .factory('apiService', apiService);

    /* @ngInject */
    function apiService($http, $location, $q, exception, logger) {
        var isPrimed = false;
        var primePromise;

        var service = {
            getContainers: getContainers,
            getContainer: getContainer,
            ready: ready
        };

        return service;

        function getContainers() {
            return $http.get('/api/containers')
                .then(getContainersComplete)
                .catch(function(message) {
                    exception.catcher('XHR Failed for getContainers')(message);
                    $location.url('/');
                });

            function getContainersComplete(data, status, headers, config) {
                return data.data[0].data.results;
            }
        }

        function prime() {
            // This function can only be called once.
            if (primePromise) {
                return primePromise;
            }

            primePromise = $q.when(true).then(success);
            return primePromise;

            function success() {
                isPrimed = true;
                logger.info('Primed data');
            }
        }

        function ready(nextPromises) {
            var readyPromise = primePromise || prime();

            return readyPromise
                .then(function() { return $q.all(nextPromises); })
                .catch(exception.catcher('"ready" function failed'));
        }

    }
})();