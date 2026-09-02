window.addEventListener("load", function load(event){

    const results_el = document.querySelector("#results");
    const submit_el = document.querySelector("#add-georeference");

    const table_el = document.createElement("table")
    table_el.setAttribute("class", "table");

    const label_th = document.createElement("th");
    label_th.appendChild(document.createTextNode("Label"));

    const id_th = document.createElement("th");
    id_th.appendChild(document.createTextNode("Identifier"));
    
    const table_head = document.createElement("tr");
    table_head.appendChild(label_th);
    table_head.appendChild(id_th);    

    table_el.appendChild(table_head);
    
    const on_select = function(label, place_id){

        return new Promise((resolve, reject) => {

	    try {
		const row = document.createElement("tr");
		
		const label_col = document.createElement("td");
		label_col.appendChild(document.createTextNode(label));
		
		const id_col = document.createElement("td");
		id_col.appendChild(document.createTextNode(place_id));
		
		row.appendChild(label_col);
		row.appendChild(id_col);
		
		table_el.appendChild(row);
		
		if (results.children.length == 0) {
		    results.appendChild(table_el);
		}
		
		resolve();
	    } catch(err) {
		reject(err);
	    }
	    
	});

    };
    
    submit_el.onclick = function(){

	sfomuseum.geocoder.georeference.init(on_select);
	return false;
    };
});
