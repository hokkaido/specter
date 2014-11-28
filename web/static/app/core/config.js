(function() {
    'use strict';

    var core = angular.module('specter.core');

    var config = {
        appErrorPrefix: '[NG-Specter Error] ', //Configure the exceptionHandler decorator
        appTitle: 'Specter',
        version: '0.0.1-SNAPSHOT'
    };

    core.value('config', config);

    core.config(configure);

    /* @ngInject */
    function configure ($logProvider, $routeProvider, routehelperConfigProvider, exceptionHandlerProvider) {
        // turn debugging off/on (no info or warn)
        if ($logProvider.debugEnabled) {
            $logProvider.debugEnabled(true);
        }

        // Configure the common route provider
        routehelperConfigProvider.config.$routeProvider = $routeProvider;
        routehelperConfigProvider.config.docTitle = 'NG-Specter: ';
        var resolveAlways = { /* @ngInject */
            ready: function(apiService) {
                return apiService.ready();
            }
            // ready: ['dataservice', function (dataservice) {
            //    return dataservice.ready();
            // }]
        };
        routehelperConfigProvider.config.resolveAlways = resolveAlways;

        // Configure the common exception handler
        exceptionHandlerProvider.configure(config.appErrorPrefix);
    }
})();