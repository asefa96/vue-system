<template>
  <span>
    <div class="info">
      <nav-bar />

 <table>
  <tr>
    <td>Host Name</td>
    <td>Smith</td>

  </tr>
  <tr>
    <td>Host IP</td>
    <td>Jackson</td>
 
  </tr>
  <tr>
    <td>System OS</td>
    <td>Johnson</td>
   
  </tr>
   <tr>
    <td>Processors</td>
    <td>Johnson</td>
   
  </tr>
   <tr>
    <td>RAM</td>
    <td>Johnson</td>
   
  </tr>
   <tr>
    <td>Disk</td>
    <td>Johnson</td>
   
  </tr>
</table>
<div class="infoCPU"> </div>
<div class="infoMemory"> </div>
    </div>
  </span>
</template>

<script>
import navBar from "./navBar.vue";
export default {
  components: {
    navBar
  }
};
</script>

<style></style>
