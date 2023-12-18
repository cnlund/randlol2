nombrechamp = document.getElementById('textchamp');
imagenchamp = document.getElementById('champimg');
btodos = document.getElementById("all");
basesinos = document.getElementById("asesino");
bpeleadores = document.getElementById("peleador");
bmagos = document.getElementById("mago");
btiradores = document.getElementById("tirador");
bsoportes = document.getElementById("soporte");
btanques = document.getElementById("tanque");

fetch("/listachamps")
    .then(response => response.json())
    .then(data => {
        let listaallchamps = {};
        for (const item of data) {
            listaallchamps[item.id] = item.tags;
        }
        function randomizar(lista) {
            const campeon = Math.floor(Math.random() * lista.length);
            return campeon;
        }
        function ponerimagen(campeon, tipos) {
            nombrechamp.innerHTML = campeon + "</br>" + tipos;
            champimg.src = `https://ddragon.leagueoflegends.com/cdn/img/champion/splash/${campeon}_0.jpg`
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
            const champelegido = randomizar(Object.keys(listaallchamps));
            ponerimagen(Object.keys(listaallchamps)[champelegido], Object.values(listaallchamps)[champelegido]);
        }

        const randasesino = function () {
             let lasesinos = filtrarchamps("Assassin")
             const asesinoelegido = randomizar(Object.keys(lasesinos));
             ponerimagen(Object.keys(lasesinos)[asesinoelegido], Object.values(lasesinos)[asesinoelegido]);
        }
        const randpeleador = function () {
            let lpeleadores = filtrarchamps("Fighter")
            const peleadorelegido = randomizar(Object.keys(lpeleadores));
            ponerimagen(Object.keys(lpeleadores)[peleadorelegido], Object.values(lpeleadores)[peleadorelegido]);
        }
        const randmago = function () {
            let lmagos = filtrarchamps("Mage")
            const magoelegido = randomizar(Object.keys(lmagos));
            ponerimagen(Object.keys(lmagos)[magoelegido], Object.values(lmagos)[magoelegido])
        }
        const randtirador = function () {
            let ltiradores = filtrarchamps("Marksman")
            const tiradorelegido = randomizar(Object.keys(ltiradores));
            ponerimagen(Object.keys(ltiradores)[tiradorelegido], Object.values(ltiradores)[tiradorelegido])
        }
        const randsoporte = function () {
            let lsoportes = filtrarchamps("Support")
            const soporteelegido = randomizar(Object.keys(lsoportes));
            ponerimagen(Object.keys(lsoportes)[soporteelegido], Object.values(lsoportes)[soporteelegido])
        }
        const randtanque = function () {
            let ltanques = filtrarchamps("Tank")
            const tanqueelegido = randomizar(Object.keys(ltanques));
            ponerimagen(Object.keys(ltanques)[tanqueelegido], Object.values(ltanques)[tanqueelegido])
        }

        btodos.addEventListener("click", randtodos)
        basesinos.addEventListener("click", randasesino)
        bpeleadores.addEventListener("click", randpeleador)
        bmagos.addEventListener("click", randmago)
        btiradores.addEventListener("click", randtirador)
        bsoportes.addEventListener("click", randsoporte)
        btanques.addEventListener("click", randtanque)
    })