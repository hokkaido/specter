(function() {
    'use strict';

    angular
        .module('specter.containers')
        .run(appRun);

    // appRun.$inject = ['routehelper']

    /* @ngInject */
    function appRun(routehelper) {
        routehelper.configureRoutes(getRoutes());
    }

    function getRoutes() {
        return [
            {
                url: '/containers',
                config: {
                    templateUrl: 'app/containers/containers.html',
                    controller: 'Containers',
                    controllerAs: 'vm',
                    title: 'containers',
                    settings: {
                        nav: 2,
                        content: '<i class="fa fa-lock"></i> Containers'
                    }
                }
            }
        ];
    }
})();