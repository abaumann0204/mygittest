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
    this.resource('people');
    this.resource('person',{path: 'people/:person_id'});
    this.route('dummy');
});


App.DummyRoute = Ember.Route.extend({
	renderTemplate: function(){	
		this.render('index',{outlet: 'index'});
		this.render('dummy',{outlet: 'dummy'});
	}
}); 

App.IndexRoute = Ember.Route.extend({
	renderTemplate: function(){	
		this.render('index',{outlet: 'index'});
	}
});

App.PersonRoute = Ember.Route.extend({

	model: function(params){
        console.log('PersonRoute: set Model ==> perform App.Person.find(params.person_id)')
        var p = App.Person.find(params.person_id);
		return p;
	},
	renderTemplate: function(){
		this.render('index',{outlet: 'index'});
		this.render('people',{outlet: 'personenliste'});
		this.render('person',{outlet: 'detailPersonData'});
	}
    
});


App.PeopleRoute = Ember.Route.extend({

	model: function(params){
        console.log('PeopleRoute: set Model ==> perform App.Person.find()')
        var people = App.Person.find();
		return people;
	},
	renderTemplate: function(){
		this.render('index',{outlet: 'index'});
		this.render('people',{outlet: 'personenliste'});
	}

});




