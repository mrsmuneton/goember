(function() {

var App = window.App = Ember.Application.create();

App.Kitten = DS.Model.extend({
  name: DS.attr('string'),
  picture: DS.attr('string')
});

DS.RESTAdapter.reopen({
  namespace: 'api/'
});

App.store = DS.Store.create({
	adapter: 'App.ApplicationAdapter'
});

App.IndexRoute = Ember.Route.extend({
	model : function(){
		return this.store.find('kitten');
	}
});

})();
