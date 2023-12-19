fetch("/listachamps")
    .then(response => response.json())
    .then(data => {
        function enviarlistachamps() {
            let listaallchamps = {};
            for (const item of data) {
                listaallchamps[item.id] = item.tags;
            }
            return listaallchamps;
        }
    })

fetch("/listachamps", { method: "POST" })
    .then(response => response.json())
    .then(data => {
        let listaallchamps = {};
        for (const item of data) {
            listaallchamps[item.id] = item.tags;
        }
        function filtrarchamps(tagespecifico) {
            listaafiltrar = Object.keys(listaallchamps)
            listafiltrada = {}
            for (let i = 0; i < listaafiltrar.length; i++) {
                let ltags = listaallchamps[listaafiltrar[i]]
                for (let j = 0; j < ltags.length; j++) {
                    if (ltags[j] === tagespecifico) {
                        listafiltrada[listaafiltrar[i]] = listaallchamps[listaafiltrar[i]];
                    }
                }
            }
            return listafiltrada
        }

        const randtodos = function () {
            listaallchamps
        }
        const randasesino = function () {
            let lasesinos = filtrarchamps("Assassin")
        }
        const randpeleador = function () {
            let lpeleadores = filtrarchamps("Fighter")
        }
        const randmago = function () {
            let lmagos = filtrarchamps("Mage")
        }
        const randtirador = function () {
            let ltiradores = filtrarchamps("Marksman")
        }
        const randsoporte = function () {
            let lsoportes = filtrarchamps("Support")
        }
        const randtanque = function () {
            let ltanques = filtrarchamps("Tank")
        }
    })