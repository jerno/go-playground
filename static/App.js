export default {
  data() {
    return {
      title: "Vue Zoo",
      learnMoreLink:
        "https://vuejs.org/guide/quick-start.html#without-build-tools",
      cards: [],
    };
  },
  methods: {
    readMore(card) {
      card.readMoreOpen = !card.readMoreOpen;
    },
    async refresh() {
      const response = await fetch("http://localhost:8081/animals/refresh");
      const data = await response.json();
      this.cards = data;
    },
  },
  async created() {
    const response = await fetch("http://localhost:8081/animals");
    const data = await response.json();
    this.cards = data;
  },
  template: `
      <top-bar :title="title"></top-bar>
      <button @click="refresh">Refresh</button>

      <div class="content">
        <div
          class="card"
          :style="{'background-image': 'url(' + card.image_link + ')'}"
          v-for="card in cards"
          v-cloak
        >
          <div
            class="scroller"
          :style="{'background-image': 'url(' + card.image_link + ')'}"
          >
            <img class="image-placeholder" :src="card.image_link" />
            <div class="text-elements">
              <h1>{{ card.name }}</h1>
              <h2><i>{{ card.latin_name }}</i></h2>

              <h3>Type:</h3>
              <p>{{ card.animal_type }}</p>
              <h3>Habitat:</h3>
              <p>{{ card.habitat }}</p>
              <h3>Diet:</h3>
              <p>{{ card.diet }}</p>
              <h3>Geo range:</h3>
              <p>{{ card.geo_range }}</p>

              <div v-if="card.readMoreOpen" class="read-more-box">
                <h3>Lifespan:</h3>
                <p>{{ card.lifespan }} years</p>
                <h3>Active time:</h3>
                <p>{{ card.active_time }}</p>
                <h3>Length:</h3>
                <p>{{ card.length_min }}-{{ card.length_max }}ft</p>
                <h3>Weight:</h3>
                <p>{{ card.weight_min }}-{{ card.weight_max }}lbs</p>
              </div>
              <p class="read-more" @click="readMore(card)">
                {{ card.readMoreOpen ? 'Read less...' : 'Read more...'}}
              </p>
            </div>
          </div>
        </div>
      </div>

      <a :href="learnMoreLink">Learn more about Vue.js</a>
  `
}