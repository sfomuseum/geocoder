'use strict';

(function(L) {
    
    if (typeof L === 'undefined') {
        throw new Error('Leaflet must be included first');
    }

    /**
     * Leaflet control that opens the geotag dialog when clicked.
     * @class
     * @extends L.Control
     * @param {Object} options
     * @param {string} [options.position='topright'] - Control position.
     * @param {function} [options.on_change] - Callback for change events.
     * @param {function} [options.on_select] - Callback for select events.
     */
    L.Control.Geocoder = L.Control.extend({
	options: {
	    position: 'topright',
	    on_change: null,
	    on_select: null,
	},

	/**
         * Called when the control is added to the map.
         * @param {L.Map} map
         * @returns {HTMLElement} Container element.
         */
	onAdd: function(map) {
	    
            var container = L.DomUtil.create('div', 'leaflet-control-geocoder leaflet-bar leaflet-control');
	    
            var link = L.DomUtil.create('a', 'leaflet-control-geocoder-button leaflet-bar-part', container);
            link.href = '#';

	    // https://icons.getbootstrap.com/icons/search/
	    link.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" fill="currentColor" class="bi bi-search" viewBox="-6 -7 30 30"><path d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001q.044.06.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1 1 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0"/></svg>';
	    
	    var icon = L.DomUtil.create('div', 'leaflet-control-geocoder-icon', link);
	    
	    this.link = link;
	    this.icon = icon;
	    this.container = container;
	    
            L.DomEvent.on(this.link, 'click', this._click, this);
	    
	    L.DomEvent.disableClickPropagation(container);
	    return container;
	},
	
	'hide': function() {
	    this.container.style.display = "none";
	},
	
	'show': function() {
	    this.container.style.display = "block";
	},
	
	onRemove: function(map) {
	    // 
	},
	
	_click: function (e) {
            L.DomEvent.stopPropagation(e);
            L.DomEvent.preventDefault(e);

	    sfomuseum.geocoder.geotag.init(this.options.on_select);
	},
    });
    
    L.Control.Geocoder.include(L.Evented.prototype);

    L.control.geocoder = function (options) {
        return new L.Control.Geocoder(options);
    };
    
})(L);
