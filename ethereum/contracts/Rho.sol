pragma solidity ^0.4.15;

contract Rho {
  function run(uint256 x0, uint256 p, uint256 n) public pure returns(uint256) {
    uint256 a = x0;
    uint256 b = x0;
    uint256 g = 1;
    while (g == 1) {
      a = a*a + p;
      b = b*b + p;
      b = b*b + p;
      if (a > b) {
        g = gcd(a - b, n);
      } else {
        g = gcd(b - a, n);
      }
    }
    return g;
  }

  function gcd(uint256 a, uint256 b) public pure returns(uint256) {
    while (b > 0) {
      a %= b;
      (a, b) = (b, a);
    }
    return a;
  }
}
