fetch("/randitems/listaitems")
    .then(response => response.json())
    .then(data => {
        listaallitems = {};
        for (const item of data){
            listaallitems[item.id] = item.name;
        }
    })