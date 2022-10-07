# visual-cryptography
Visual Cryptography is a cryptographic technique which allows visual image to be encrypted in such a way that the decrypted information is a visual image. One of the best-known schemes is by Adi Shamir & Moni Naor developed 1994. ![[paper]](https://link.springer.com/content/pdf/10.1007/BFb0053419.pdf) <br>
visual cryptography is easy to decode, implement and requires no cryptographic computations. It can be used for secret sharing, In which an  image is broken into n shares and only someone with all the n shares is able to decrypt the image. The method can be extended to allow k out of n shares to be enough for the secret to be revealed.<br>
Here I implemented n out of n scheme for small values of n. To understand how everything works, I recommend reading the original paper. Alternatively you can read [this](https://datagenetics.com/blog/november32013/index.html) blog post by Nick Berry. There are also comments and implementation notes in the source code.

# Usage

# Examples
<br>
<p align="center">
  <img src="assets/example_2_2/img_0.png"> 
    <img src="assets/example_2_2/img_1.png"> 
      <img src="assets/example_2_2/merged.png"> 
</p>
