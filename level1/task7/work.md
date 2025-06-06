# ABI 编码中的 uint<M> 表示什么？
在 ABI 编码中，uint<M> 表示 M 位的无符号整数，其中 M 是一个 8 的倍数，且 0 < M <= 256。
`uint<M>` 类型的数据会被编码为 32 字节（256 位）的固定长度。若 `M` 小于 256，数据会在左侧用零填充到 32 字节。例如，`uint8` 类型的数值 `42` 编码后是 `0x000000000000000000000000000000000000000000000000000000000000002A`。
# 在 ABI 中，动态类型和静态类型有什么区别？
静态类型的大小和位置在编码前是已知的，可以直接编码。动态类型的大小或位置在编码前可能不确定，需要在编码后的数据中单独指定。
# 解释函数选择器(function selector)在 ABI 中的用途。
函数选择器是一个 4 字节（32 位）的哈希值，用于唯一标识智能合约中的某个函数。它通过对函数签名（函数名和参数类型）进行 `keccak256` 哈希计算，然后取前 4 个字节得到。例如，对于函数 `transfer(address,uint256)`，其函数选择器计算方式如下：
# 在 Solidity 中，哪些类型不被 ABI 直接支持？
Solidity 中的元组类型不被 ABI 直接支持，需要特定的处理。

# 如何通过 ABI 编码调用具有多个参数的函数？
通过将所有参数的编码合并，其中静态参数直接编码，动态参数先记录偏移量然后在数据部分单独编码。

# 什么是“严格编码模式”？
严格编码模式要求编码偏移量必须尽可能小，且数据区域不允许有重叠或间隙。

# 在 ABI 中，fixed<M>x<N> 和 ufixed<M>x<N> 有何不同？
fixed<M>x<N> 是有符号的固定小数点数，而 ufixed<M>x<N> 是无符号的固定小数点数。其中 M 是总位数，N 是小数位数。

# 事件的 ABI 编码如何处理已索引和未索引的参数？
已索引的参数将与事件的 Keccak 哈希一起作为日志项的主题存储。未索引的参数则存储在日志的数据部分。

# 描述如何通过 ABI 对一个返回错误的函数进行编码。
错误函数的编码与普通函数相似，但使用错误选择器。例如，InsufficientBalance 错误将编码其参数并使用特定的错误选择器。

# abi.encodePacked() 在什么情况下使用，它与 abi.encode() 有何区别？
abi.encodePacked() 用于非标准打包模式，适用于需要紧凑编码的情况。它与 abi.encode() 的主要区别是不会对短于 32 字节的类型进行补 0 操作，且动态类型不包含长度信息。

# 解释 ABI 中对动态数组编码的过程。
动态数组首先编码数组长度，然后编码数组中每个元素。如果元素是动态类型，则对每个元素进行独立编码，并记录其偏移。

# 如何在 ABI 中处理嵌套数组或结构体？
嵌套数组或结构体按其元素顺序编码，每个元素根据其类型（静态或动态）适当处理。动态元素会记录偏移量，然后编码其内容。