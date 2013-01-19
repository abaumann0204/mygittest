/*

		Version unter Nutzung von "nested routing"

*/

window.App = Ember.Application.create();

DS.RESTAdapter.configure("plurals",{"person" : "people"})

App.Store = DS.Store.extend({
	revision: 11,
    adapter: DS.RESTAdapter.create({namespace: 'restservice'})
});

var attr = DS.attr;

App.Person = DS.Model.extend({
	firstName: attr('string'),
	lastName: attr('string'),
	birthDay: attr('string')
});



App.Router.map(function(){

    	 this.resource('people',function(){
    	        this.resource('person',{path: 'people/:person_id'});
        });

    	this.route('page1');
        this.route('page2');
});


App.IndexRoute = Ember.Route.extend({

});

App.Page1Route = Ember.Route.extend({

});


App.Page2Route = Ember.Route.extend({

});


App.PersonRoute = Ember.Route.extend({

	model: function(params){
        console.log('PersonRoute: set Model ==> perform App.Person.find(params.person_id)')
        var p = App.Person.find(params.person_id);
		return p;
	}
    
});


App.PeopleRoute = Ember.Route.extend({

	model: function(params){
        console.log('PeopleRoute: set Model ==> perform App.Person.find()')
        var people = App.Person.find();
		return people;
	}

});




