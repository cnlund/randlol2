nombrechamp = document.getElementById('textchamp');
imagenchamp = document.getElementById('champimg');
btodos = document.getElementById("all");

fetch("/listachamps")
    .then(response => response.json())
    .then(data => {
        let listaallchamps = {};
        for (const item of data) {
            listaallchamps[item.id] = item.tags; 
        }
        function randomizar(lista){
            const campeon = Math.floor(Math.random() * lista.length);
            return campeon;
        }
        function ponerimagen(campeon){
            nombrechamp.innerHTML = campeon;
            champimg.src = `https://ddragon.leagueoflegends.com/cdn/img/champion/splash/${campeon}_0.jpg`
            champimg.style.display = flex;
        }
        
        const randtodos = function(){
            const champelegido = randomizar(Object.keys(listaallchamps));
            ponerimagen(Object.keys(listaallchamps)[champelegido]);  
        }

        btodos.addEventListener("click", randtodos)
    })