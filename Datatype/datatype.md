# Data Type in GoLang

+ ## Boolean

  + ### true
  + ### false

```go
func main(){
    n := 1 == 1
    m := 1 == 2 
    fmt.Printf("%v, %T", n, n)  // true, bool
    fmt.Printf("%v, %T", m, m)  // false, bool

}
```

### If **boolean variable** is initialized without any value, it's default value is *false*

```go
var n bool
fmt.Printf("%v, %T", n, n)  // Output: false, bool
```

+ ## Numeric Datatype

<details>
        <summary>Integer-type</summary>
        <dl>
            <dt>
                <h3>uint8</h3>
            </dt>
            <h4>
                <dd>Unsigned 8-bit integers (0 to 255)</dd>
            </h4>
            <dt>
                <h3>uint16</h3>
            </dt>
            <h4>
                <dd>Unsigned 16-bit integers (0 to 65,535)</dd>
            <dt>
            <dt>
                <h3>uint32</h3>
            </dt>
            <h4>
                <dd>Unsigned 32-bit integers (0 to 4,29,49,67,295)</dd>
            <dt>
            <dt>
                <h3>uint64</h3>
            </dt>
            <h4>
                <dd>Unsigned 64-bit integers (0 to 1,84,46,74,40,73,70,95,51,615)</dd>
            <dt>
                <h3>int8</h3>
            </dt>
            <h4>
                <dd> 8 bit Signed Integer (two’s complement)</dd>
            </h4>
            <h4>
                <dd><b>Range: </b> -128 to 127</dd>
            </h4>
            <dt>
                <h3>int16</h3>
            </dt>
            <h4>
                <dd> 16 bit Signed Integer (two’s complement)</dd>
            </h4>
            <h4>
                <dd><b>Range: </b> -2<sup>15</sup> to 2<sup>15</sup> -1 :: (-32,768 to 32,767)</dd>
            </h4>
            <dt>
                <h3>int32</h3>
            </dt>
            <h4>
                <dd> 32 bit Signed Integer (two’s complement)</dt>
            </h4>
            <h4>
                <dd><b>Range: </b> -2<sup>31</sup> to 2<sup>31</sup> -1 :: (-2,14,74,83,648 to 2,14,74,83,647)</dd>
            </h4>
            <dt>
                <h3>int64</h3>
            </dt>
            <h4>
                <dd> 64 bit Signed Integer (two’s complement). They can also be represented in octal and hexadecimal
                    </dt>
            </h4>
            <h4>
                <dd><b>Range: </b> -2<sup>63</sup> to 2<sup>63</sup> -1 :: (-92,23,37,20,36,85,47,75,808 to
                    92,23,37,20,36,85,47,75,807)</dd>
            </h4>
            <dt>
                <h3>int</h3>
            </dt>
            <h4>
                <dd> Signed integers of at least 32-bit in size; not equivalent to int32. It is 32 bits wide on a 32-bit
                    system and 64-bits wide on a 64-bit system</dt>
            </h4>
            <h4>
                <dd><b>Range: </b> Platform dependent</dd>
            </h4>
        </dl>
    </details>
    <details>
        <summary>Float-type</summary>
        <dl>
            <dt>
                <h3>float32</h3>
            </dt>
            <h4>
                <dd>IEEE-754 32-bit floating-point numbers</dd>
            </h4>
            <dt>
                <h3>float64</h3>
            </dt>
            <h4>
                <dd>IEEE-754 64-bit floating-point numbers</dd>
            </h4>
            <dt>
                <h3>complex64</h3>
            </dt>
            <h4>
                <dd>Complex numbers with float32 real and imaginary parts</dd>
            </h4>
            <dt>
                <h3>complex128</h3>
            </dt>
            <h4>
                <dd>Complex numbers with float64 real and imaginary parts</dd>
            </h4>
        </dl>
    </details>
    <details>
        <summary>Other Numeric-type</summary>
        <dl>
            <dt>
                <h3>byte</h3>
            </dt>
            <h4>
                <dd>same as uint8</dd>
            </h4>
            <dt>
                <h3>rune</h3>
            </dt>
            <h4>
                <dd>same as int32</dd>
            </h4>
            <dt>
                <h3>uint</h3>
            </dt>
            <h4>
                <dd>32 or 64 bits</dd>
            </h4>
            <dt>
                <h3>int</h3>
            </dt>
            <h4>
                <dd>same size as uint</dd>
            </h4>
            <dt>
                <h3>uintptr</h3>
            </dt>
            <h4>
                <dd>an unsigned integer to store the uninterpreted bits of a pointer value</dd>
            </h4>
        </dl>
    </details>

+ ## String

```go
// Strings in GoLang
s := "A String"  // double-quotes for Strings
b := []byte(s)  // converted to byte >> would output ASCII values of characters
fmt.Printf("%v, %T", b, b)  // Output: [65 32 83 116 114 105 110 103], []uint8

// Characters in GoLang
r := 'a'  // single-quotes for Character  # same-as: var r rune = 'a'
fmt.Printf("%v, %T", r, r)  // Output: 97, int32  # ASCII value of 'a'
```

<!-- + ## Derived -->

## Summary

+ ## Boolean
  + ### Values are **true** or **false**
	+ ### Zero value is **false** as well as **default**

+ ## Numeric

  + ## Integer

	  + ### Signed Integer
		  + ### int type has varying size, but minium 32 bits
		  + ### Range: 8 bit (int 8) to 64 bit (int64)
	  + ### Unsigned Integer
		  + ### Range: 8 bit (byte and uint8) to 32 bit (uint32)
	  + ### Arithmetic operations can be performed on same data-type ({int to int}, {float32 to float32})
	  + ### Can't mix types in same family **(uint16 + uint32 = error)**

  + ## Float
	  + ### Follows **IEEE-754 standard**
		+ ### 32 bit and 64 bit versions
		+ ### Literal Style
		  + ### Decimal (3.14)
		  + ### Exponential (13e18 or 2E10)
		  + ### Mixed (13.7e12)

  + ## Complex
    + ### Zero Value (0+0i)
    + ### 64 bit and 128 bit versions
    + ### Built-in functions
      + ### **complex** : make complex number from two floats
      + ### **real** : get real part as float
      + ### **imag** : get imaginary part as float

  + ## Text Type
	  + ## Strings
		  + ### UTF-8
		  + ### Immutable
		  + ### Can be concatenated with **plus(+)** operator
		  + ### can be converted to **[ ]byte**
	  + ## Rune
		  + ### UTF-32
		  + ### Alias for **int32**
		  + ### Special methods normally required to process
		    + ### e.g.: strings.Reader#ReadRunet
