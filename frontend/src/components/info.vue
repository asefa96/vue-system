<template>
  <span>
    <div class="info">
      <nav-bar />

 <table>
  <tr>
    <td><fa icon="home" fixed-width /> Host Name</td>
    <td >  {{data.hostname}}</td>

  </tr>
  <tr>
    <td><fa icon="network-wired" fixed-width /> Host IP</td>
    <td>{{data.hostip}}</td>
 
  </tr>
  <tr>
    <td><fa icon="cogs" fixed-width /> System OS</td>
    <td>{{data.systemos}}</td>
   
  </tr>
   <tr>
    <td><fa icon="microchip" fixed-width /> Processors</td>
    <td> {{data.hostname}} {{data.processors}}</td>
   
  </tr>
   <tr>
    <td><fa icon="memory" fixed-width /> RAM</td>
    <td>{{data.ram}} GB</td>
   
  </tr>
   <tr>
    <td><fa icon="server" fixed-width /> Disk Used</td>
    <td>% {{parseFloat(data.diskused).toFixed(3)}}</td>
   
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
  props:['data'],
  components: {
    navBar
  }
};
</script>

<style></style>
