package wstat_test

import (
	"fmt"
	"strings"

	"github.com/mdigger/wstat"
)

func ExampleFromString() {
	sample := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Nunc sit amet ipsum vel nunc interdum ultricies eu non augue. Donec sit amet 
nisl aliquet, ultricies enim id, malesuada libero. Ut maximus felis neque, sed 
porta est accumsan in. Curabitur tincidunt fringilla ultrices. Suspendisse 
porttitor non mauris quis tincidunt. Vivamus sit amet ante vel dui pellentesque 
mollis sit amet sit amet nibh. Vivamus id ante ultricies mi tincidunt sodales et 
pretium ex. In placerat purus vitae ligula tincidunt consectetur. Vivamus vel 
leo ut est molestie molestie non et odio. Nam a iaculis magna, sit amet accumsan 
elit. Nullam quam sapien, accumsan nec porta non, sollicitudin ut magna. 
Suspendisse sed gravida nisl. Nullam porta ultricies pellentesque. Nunc viverra 
convallis mauris, ac aliquam velit commodo in. Nulla facilisis commodo massa in 
egestas. Quisque at enim risus.

Nulla facilisi. Morbi odio ligula, hendrerit vitae mi ullamcorper, fermentum 
laoreet ligula. Aliquam ornare enim nec tortor sagittis faucibus. Morbi pretium 
dui at nibh placerat semper. Maecenas et libero vitae orci fringilla pretium sit 
amet a est. In at ipsum est. Sed laoreet efficitur consequat. Ut pharetra mauris 
sed mi consequat, ac suscipit dolor convallis. Vestibulum in est sollicitudin, 
mattis urna a, malesuada felis. Duis nibh lectus, viverra in aliquet sed, 
ullamcorper et justo. In et elementum sem.
	
Vivamus purus tellus, feugiat ac convallis sed, sollicitudin id justo. Donec 
aliquam ullamcorper ipsum, congue pretium dui interdum a. Maecenas vel neque ac 
magna ornare tempus. Pellentesque tincidunt tincidunt sollicitudin. Morbi neque 
nulla, porttitor vel sagittis quis, dapibus ut leo. In a arcu nec magna cursus 
porta. Donec fermentum dolor a augue viverra feugiat vel eu odio. Sed eu dapibus 
libero. Quisque lacus risus, accumsan ac suscipit non, molestie vel neque. 
Aliquam consequat non neque at molestie. Nunc sed erat ultrices, viverra elit 
quis, tincidunt purus. Fusce vitae diam auctor, ultricies massa at, dictum metus. 
Ut at nibh id velit sollicitudin facilisis ut sit amet dui. Sed ac sapien 
dignissim, accumsan metus et, tempor est.
	
Praesent mollis sagittis neque vel pellentesque. Phasellus laoreet sollicitudin 
ante quis consectetur. Pellentesque hendrerit porta commodo. Proin eget congue 
mauris. Ut nec ornare tellus, id rhoncus nibh. Donec eget elit non nunc egestas 
tempor ac quis massa. Ut nisi augue, gravida in quam aliquet, mollis varius 
augue. Nam vehicula commodo egestas. Phasellus vel odio sollicitudin, sodales 
lacus non, lobortis lorem. Quisque nisl metus, porta vitae mollis sit amet, 
semper eu nulla. Maecenas rhoncus urna ac lacus facilisis, fringilla suscipit 
libero pharetra. Aliquam ornare metus eget magna accumsan tincidunt.
	
In pellentesque neque vel ex sodales feugiat vel nec nibh. Nullam eleifend velit 
at enim congue tempor. Suspendisse gravida gravida enim id convallis. Class 
aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos 
himenaeos. Pellentesque rutrum orci in mi consectetur, sit amet bibendum elit 
vehicula. Curabitur tincidunt metus id ex pulvinar, in cursus dui interdum. 
Curabitur nec semper lectus, a tempor dui. Integer porta ligula nec sollicitudin 
feugiat. Vivamus ornare ligula vel elit pellentesque, id sagittis neque 
dignissim. Nulla purus nunc, fermentum ut efficitur eu, ultrices vel odio. 
Proin id accumsan nisi. Praesent sem felis, lacinia vel quam a, interdum 
fringilla velit.

1234567890.`

	stat := wstat.FromString(sample)
	fmt.Printf(`
--- stats -----------
chars:        %v
spaces:       %v
puncts:       %v
numbers:      %v
words:        %v
--- pages -----------
typewritten:  %v
author's:     %v
--- reading time ----
duration:     %v
---------------------
`,
		stat.Chars, stat.Spaces, stat.Puncts, stat.Numbers, stat.Words,
		stat.Pages(), stat.AuthorPages(), stat.Duration(228))

	// Output:
	// --- stats -----------
	// chars:        3504
	// spaces:       566
	// puncts:       111
	// numbers:      0
	// words:        518
	// --- pages -----------
	// typewritten:  2
	// author's:     0.0876
	// --- reading time ----
	// duration:     2m16s
	// ---------------------
}

func ExampleFromHTML() {
	sample := `<html>
<head>
<title>   123    </title>
<style type="text/css">
<!--
h1	{text-align:center;
	font-family:Arial, Helvetica, Sans-Serif;
	}

p	{text-indent:20px;
	}
-->
</style>
</head>
<body bgcolor = "#ffffcc" text = "#000000">
<h1>Lorem Ipsum</h1>

<p><strong>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</strong><br/>
Nunc sit amet ipsum vel nunc interdum ultricies eu non augue. Donec sit amet 
nisl aliquet, ultricies enim id, malesuada libero. Ut maximus felis neque, sed 
porta est accumsan in. Curabitur tincidunt fringilla ultrices. Suspendisse 
porttitor non mauris quis tincidunt. Vivamus sit amet ante vel dui pellentesque 
mollis sit amet sit amet nibh. Vivamus id ante ultricies mi tincidunt sodales et 
pretium ex. In placerat purus vitae ligula tincidunt consectetur. Vivamus vel 
leo ut est molestie molestie non et odio. Nam a iaculis magna, sit amet accumsan 
elit. Nullam quam sapien, accumsan nec porta non, sollicitudin ut magna. 
Suspendisse sed gravida nisl. Nullam porta ultricies pellentesque. Nunc viverra 
convallis mauris, ac aliquam velit commodo in. Nulla facilisis commodo massa in 
egestas. Quisque at enim risus.</p>

<p>Nulla facilisi. Morbi odio ligula, hendrerit vitae mi ullamcorper, fermentum 
laoreet ligula. Aliquam ornare enim nec tortor sagittis faucibus. Morbi pretium 
dui at nibh placerat semper. Maecenas et libero vitae orci fringilla pretium sit 
amet a est. In at ipsum est. Sed laoreet efficitur consequat. Ut pharetra mauris 
sed mi consequat, ac suscipit dolor convallis. Vestibulum in est sollicitudin, 
mattis urna a, malesuada felis. Duis nibh lectus, viverra in aliquet sed, 
ullamcorper et justo. In et elementum sem.</p>
	
<p>Vivamus purus tellus, feugiat ac convallis sed, sollicitudin id justo. Donec 
aliquam ullamcorper ipsum, congue pretium dui interdum a. Maecenas vel neque ac 
magna ornare tempus. Pellentesque tincidunt tincidunt sollicitudin. Morbi neque 
nulla, porttitor vel sagittis quis, dapibus ut leo. In a arcu nec magna cursus 
porta. Donec fermentum dolor a augue viverra feugiat vel eu odio. Sed eu dapibus 
libero. Quisque lacus risus, accumsan ac suscipit non, molestie vel neque. 
Aliquam consequat non neque at molestie. Nunc sed erat ultrices, viverra elit 
quis, tincidunt purus. Fusce vitae diam auctor, ultricies massa at, dictum metus. 
Ut at nibh id velit sollicitudin facilisis ut sit amet dui. Sed ac sapien 
dignissim, accumsan metus et, tempor est.</p>
	
<p>Praesent mollis sagittis neque vel pellentesque. Phasellus laoreet sollicitudin 
ante quis consectetur. Pellentesque hendrerit porta commodo. Proin eget congue 
mauris. Ut nec ornare tellus, id rhoncus nibh. Donec eget elit non nunc egestas 
tempor ac quis massa. Ut nisi augue, gravida in quam aliquet, mollis varius 
augue. Nam vehicula commodo egestas. Phasellus vel odio sollicitudin, sodales 
lacus non, lobortis lorem. Quisque nisl metus, porta vitae mollis sit amet, 
semper eu nulla. Maecenas rhoncus urna ac lacus facilisis, fringilla suscipit 
libero pharetra. Aliquam ornare metus eget magna accumsan tincidunt.</p>
	
<p>In pellentesque neque vel ex sodales feugiat vel nec nibh. Nullam eleifend velit 
at enim congue tempor. Suspendisse gravida gravida enim id convallis. Class 
aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos 
himenaeos. Pellentesque rutrum orci in mi consectetur, sit amet bibendum elit 
vehicula. Curabitur tincidunt metus id ex pulvinar, in cursus dui interdum. 
Curabitur nec semper lectus, a tempor dui. Integer porta ligula nec sollicitudin 
feugiat. Vivamus ornare ligula vel elit pellentesque, id sagittis neque 
dignissim. Nulla purus nunc, fermentum ut efficitur eu, ultrices vel odio. 
Proin id accumsan nisi. Praesent sem felis, lacinia vel quam a, interdum 
fringilla velit.</p>

</body>
</html>`

	stat, err := wstat.FromHTML(strings.NewReader(sample))
	if err != nil {
		panic(err)
	}

	fmt.Println("reading time:", stat)

	// Output:
	// reading time: 2m36s (520 words)
}
