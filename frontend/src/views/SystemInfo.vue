<template>
  <span>
    <info :data="data" />
  </span>
</template>

<script>
import Info from "../components/info.vue";
import { apiGet } from "../api/getData";
export default {
  components: { Info },
  data() {
    return {
      data: [],
    };
  },
  async mounted() {
    this.data=await apiGet("");
    console.log(this.data)
  },
};
</script>

<style></style>
