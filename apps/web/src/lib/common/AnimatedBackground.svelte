<ul class="background">
  <li></li>
  <li></li>
  <li></li>
  <li></li>
  <li></li>
  <li></li>
  <li></li>
  <li></li>
  <li></li>
  <li></li>
</ul>

<style>
  @keyframes animate {
    0% {
      transform: translateY(0) rotate(0deg);
      opacity: 1;
      border-radius: 0;
    }
    100% {
      transform: translateY(-1000px) rotate(720deg);
      opacity: 0;
      border-radius: 50%;
    }
  }

  .background {
    z-index: -1;
    position: fixed;
    width: 100vw;
    height: 100vh;
    top: 0;
    left: 0;
    margin: 0;
    padding: 0;
    background: #02031d;
    overflow: hidden;
  }
  .background li {
    position: absolute;
    display: block;
    list-style: none;
    width: 20px;
    height: 20px;
    background: rgba(255, 255, 255, 0.1);
    animation: animate 33s linear infinite;
  }

  .background li:nth-child(0) {
    left: 56%;
    width: 39px;
    height: 39px;
    bottom: -39px;
  }
  .background li:nth-child(1) {
    left: 10%;
    width: 104px;
    height: 104px;
    bottom: -104px;
    animation-delay: 8s;
  }
  .background li:nth-child(2) {
    left: 8%;
    width: 99px;
    height: 99px;
    bottom: -99px;
    animation-delay: 12s;
  }
  .background li:nth-child(3) {
    left: 47%;
    width: 156px;
    height: 156px;
    bottom: -156px;
    animation-delay: 30s;
  }
  .background li:nth-child(4) {
    left: 48%;
    width: 153px;
    height: 153px;
    bottom: -153px;
    animation-delay: 15s;
  }
  .background li:nth-child(5) {
    left: 46%;
    width: 109px;
    height: 109px;
    bottom: -109px;
    animation-delay: 5s;
  }
  .background li:nth-child(6) {
    left: 1%;
    width: 74px;
    height: 74px;
    bottom: -74px;
    animation-delay: 19s;
  }
  .background li:nth-child(7) {
    left: 34%;
    width: 56px;
    height: 56px;
    bottom: -56px;
    animation-delay: 29s;
  }
  .background li:nth-child(8) {
    left: 36%;
    width: 123px;
    height: 123px;
    bottom: -123px;
    animation-delay: 3s;
  }
  .background li:nth-child(9) {
    left: 35%;
    width: 70px;
    height: 70px;
    bottom: -70px;
    animation-delay: 43s;
  }
</style>
