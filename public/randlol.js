fetch("/listachamps")
    .then(response => response.json())
    .then(data => {
        let listaallchamps = {};
        for (const item of data) {
            listaallchamps[item.name] = item.tags; 
        }
        console.log(listaallchamps);
    })