// Code generated from execution/strictus/parser/Strictus.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Strictus
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "strings"

var _ = strings.Builder{}

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 48, 400,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 3, 2, 7, 2,
	88, 10, 2, 12, 2, 14, 2, 91, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 97, 10,
	3, 3, 4, 5, 4, 100, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 106, 10, 4, 3,
	4, 3, 4, 3, 4, 5, 4, 111, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 7, 5, 118,
	10, 5, 12, 5, 14, 5, 121, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 7,
	7, 129, 10, 7, 12, 7, 14, 7, 132, 11, 7, 3, 8, 3, 8, 5, 8, 136, 10, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 145, 10, 9, 12, 9, 14,
	9, 148, 11, 9, 5, 9, 150, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 155, 10, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 164, 10, 11, 12, 11,
	14, 11, 167, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 175,
	10, 12, 3, 13, 3, 13, 5, 13, 179, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 5, 14, 187, 10, 14, 5, 14, 189, 10, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 199, 10, 16, 3, 16, 3, 16, 3,
	16, 3, 17, 3, 17, 7, 17, 206, 10, 17, 12, 17, 14, 17, 209, 11, 17, 3, 17,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5,
	19, 222, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 230,
	10, 20, 12, 20, 14, 20, 233, 11, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 7, 21, 241, 10, 21, 12, 21, 14, 21, 244, 11, 21, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 253, 10, 22, 12, 22, 14, 22, 256,
	11, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 265, 10,
	23, 12, 23, 14, 23, 268, 11, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 7, 24, 277, 10, 24, 12, 24, 14, 24, 280, 11, 24, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25, 289, 10, 25, 12, 25, 14, 25,
	292, 11, 25, 3, 26, 3, 26, 6, 26, 296, 10, 26, 13, 26, 14, 26, 297, 3,
	26, 3, 26, 5, 26, 302, 10, 26, 3, 27, 3, 27, 7, 27, 306, 10, 27, 12, 27,
	14, 27, 309, 11, 27, 3, 28, 3, 28, 5, 28, 313, 10, 28, 3, 29, 3, 29, 3,
	30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 5, 34, 330, 10, 34, 3, 34, 3, 34, 3, 34, 5, 34, 335, 10,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 342, 10, 34, 3, 35, 3, 35,
	5, 35, 346, 10, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 3, 38, 7, 38, 359, 10, 38, 12, 38, 14, 38, 362, 11, 38,
	5, 38, 364, 10, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 5, 39, 371, 10,
	39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 379, 10, 41, 3, 42,
	3, 42, 3, 42, 3, 42, 7, 42, 385, 10, 42, 12, 42, 14, 42, 388, 11, 42, 5,
	42, 390, 10, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 398,
	10, 43, 3, 43, 2, 8, 38, 40, 42, 44, 46, 48, 44, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 2,
	9, 3, 2, 33, 34, 3, 2, 16, 17, 3, 2, 18, 21, 3, 2, 22, 23, 3, 2, 24, 26,
	4, 2, 23, 23, 27, 27, 3, 2, 38, 39, 2, 408, 2, 89, 3, 2, 2, 2, 4, 96, 3,
	2, 2, 2, 6, 99, 3, 2, 2, 2, 8, 114, 3, 2, 2, 2, 10, 122, 3, 2, 2, 2, 12,
	126, 3, 2, 2, 2, 14, 133, 3, 2, 2, 2, 16, 154, 3, 2, 2, 2, 18, 156, 3,
	2, 2, 2, 20, 165, 3, 2, 2, 2, 22, 174, 3, 2, 2, 2, 24, 176, 3, 2, 2, 2,
	26, 180, 3, 2, 2, 2, 28, 190, 3, 2, 2, 2, 30, 194, 3, 2, 2, 2, 32, 203,
	3, 2, 2, 2, 34, 213, 3, 2, 2, 2, 36, 215, 3, 2, 2, 2, 38, 223, 3, 2, 2,
	2, 40, 234, 3, 2, 2, 2, 42, 245, 3, 2, 2, 2, 44, 257, 3, 2, 2, 2, 46, 269,
	3, 2, 2, 2, 48, 281, 3, 2, 2, 2, 50, 301, 3, 2, 2, 2, 52, 303, 3, 2, 2,
	2, 54, 312, 3, 2, 2, 2, 56, 314, 3, 2, 2, 2, 58, 316, 3, 2, 2, 2, 60, 318,
	3, 2, 2, 2, 62, 320, 3, 2, 2, 2, 64, 322, 3, 2, 2, 2, 66, 341, 3, 2, 2,
	2, 68, 345, 3, 2, 2, 2, 70, 347, 3, 2, 2, 2, 72, 350, 3, 2, 2, 2, 74, 354,
	3, 2, 2, 2, 76, 370, 3, 2, 2, 2, 78, 372, 3, 2, 2, 2, 80, 378, 3, 2, 2,
	2, 82, 380, 3, 2, 2, 2, 84, 397, 3, 2, 2, 2, 86, 88, 5, 4, 3, 2, 87, 86,
	3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2,
	90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 7, 2, 2, 3, 93, 3, 3, 2,
	2, 2, 94, 97, 5, 6, 4, 2, 95, 97, 5, 30, 16, 2, 96, 94, 3, 2, 2, 2, 96,
	95, 3, 2, 2, 2, 97, 5, 3, 2, 2, 2, 98, 100, 7, 31, 2, 2, 99, 98, 3, 2,
	2, 2, 99, 100, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 7, 30, 2, 2,
	102, 103, 7, 40, 2, 2, 103, 105, 7, 28, 2, 2, 104, 106, 5, 8, 5, 2, 105,
	104, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 110,
	7, 29, 2, 2, 108, 109, 7, 3, 2, 2, 109, 111, 5, 12, 7, 2, 110, 108, 3,
	2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 113, 5, 18, 10,
	2, 113, 7, 3, 2, 2, 2, 114, 119, 5, 10, 6, 2, 115, 116, 7, 4, 2, 2, 116,
	118, 5, 10, 6, 2, 117, 115, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117,
	3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 9, 3, 2, 2, 2, 121, 119, 3, 2, 2,
	2, 122, 123, 7, 40, 2, 2, 123, 124, 7, 5, 2, 2, 124, 125, 5, 12, 7, 2,
	125, 11, 3, 2, 2, 2, 126, 130, 5, 16, 9, 2, 127, 129, 5, 14, 8, 2, 128,
	127, 3, 2, 2, 2, 129, 132, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 131,
	3, 2, 2, 2, 131, 13, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 133, 135, 7, 6,
	2, 2, 134, 136, 7, 41, 2, 2, 135, 134, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2,
	136, 137, 3, 2, 2, 2, 137, 138, 7, 7, 2, 2, 138, 15, 3, 2, 2, 2, 139, 155,
	7, 40, 2, 2, 140, 149, 7, 28, 2, 2, 141, 146, 5, 12, 7, 2, 142, 143, 7,
	4, 2, 2, 143, 145, 5, 12, 7, 2, 144, 142, 3, 2, 2, 2, 145, 148, 3, 2, 2,
	2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148,
	146, 3, 2, 2, 2, 149, 141, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 151,
	3, 2, 2, 2, 151, 152, 7, 29, 2, 2, 152, 153, 7, 3, 2, 2, 153, 155, 5, 12,
	7, 2, 154, 139, 3, 2, 2, 2, 154, 140, 3, 2, 2, 2, 155, 17, 3, 2, 2, 2,
	156, 157, 7, 8, 2, 2, 157, 158, 5, 20, 11, 2, 158, 159, 7, 9, 2, 2, 159,
	19, 3, 2, 2, 2, 160, 161, 5, 22, 12, 2, 161, 162, 5, 84, 43, 2, 162, 164,
	3, 2, 2, 2, 163, 160, 3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165, 163, 3, 2,
	2, 2, 165, 166, 3, 2, 2, 2, 166, 21, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2,
	168, 175, 5, 24, 13, 2, 169, 175, 5, 26, 14, 2, 170, 175, 5, 28, 15, 2,
	171, 175, 5, 4, 3, 2, 172, 175, 5, 32, 17, 2, 173, 175, 5, 34, 18, 2, 174,
	168, 3, 2, 2, 2, 174, 169, 3, 2, 2, 2, 174, 170, 3, 2, 2, 2, 174, 171,
	3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 174, 173, 3, 2, 2, 2, 175, 23, 3, 2,
	2, 2, 176, 178, 7, 32, 2, 2, 177, 179, 5, 34, 18, 2, 178, 177, 3, 2, 2,
	2, 178, 179, 3, 2, 2, 2, 179, 25, 3, 2, 2, 2, 180, 181, 7, 35, 2, 2, 181,
	182, 5, 34, 18, 2, 182, 188, 5, 18, 10, 2, 183, 186, 7, 36, 2, 2, 184,
	187, 5, 26, 14, 2, 185, 187, 5, 18, 10, 2, 186, 184, 3, 2, 2, 2, 186, 185,
	3, 2, 2, 2, 187, 189, 3, 2, 2, 2, 188, 183, 3, 2, 2, 2, 188, 189, 3, 2,
	2, 2, 189, 27, 3, 2, 2, 2, 190, 191, 7, 37, 2, 2, 191, 192, 5, 34, 18,
	2, 192, 193, 5, 18, 10, 2, 193, 29, 3, 2, 2, 2, 194, 195, 9, 2, 2, 2, 195,
	198, 7, 40, 2, 2, 196, 197, 7, 5, 2, 2, 197, 199, 5, 12, 7, 2, 198, 196,
	3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 201, 7, 10,
	2, 2, 201, 202, 5, 34, 18, 2, 202, 31, 3, 2, 2, 2, 203, 207, 7, 40, 2,
	2, 204, 206, 5, 68, 35, 2, 205, 204, 3, 2, 2, 2, 206, 209, 3, 2, 2, 2,
	207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 210, 3, 2, 2, 2, 209,
	207, 3, 2, 2, 2, 210, 211, 7, 10, 2, 2, 211, 212, 5, 34, 18, 2, 212, 33,
	3, 2, 2, 2, 213, 214, 5, 36, 19, 2, 214, 35, 3, 2, 2, 2, 215, 221, 5, 38,
	20, 2, 216, 217, 7, 11, 2, 2, 217, 218, 5, 34, 18, 2, 218, 219, 7, 5, 2,
	2, 219, 220, 5, 34, 18, 2, 220, 222, 3, 2, 2, 2, 221, 216, 3, 2, 2, 2,
	221, 222, 3, 2, 2, 2, 222, 37, 3, 2, 2, 2, 223, 224, 8, 20, 1, 2, 224,
	225, 5, 40, 21, 2, 225, 231, 3, 2, 2, 2, 226, 227, 12, 3, 2, 2, 227, 228,
	7, 12, 2, 2, 228, 230, 5, 40, 21, 2, 229, 226, 3, 2, 2, 2, 230, 233, 3,
	2, 2, 2, 231, 229, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 39, 3, 2, 2,
	2, 233, 231, 3, 2, 2, 2, 234, 235, 8, 21, 1, 2, 235, 236, 5, 42, 22, 2,
	236, 242, 3, 2, 2, 2, 237, 238, 12, 3, 2, 2, 238, 239, 7, 13, 2, 2, 239,
	241, 5, 42, 22, 2, 240, 237, 3, 2, 2, 2, 241, 244, 3, 2, 2, 2, 242, 240,
	3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 41, 3, 2, 2, 2, 244, 242, 3, 2,
	2, 2, 245, 246, 8, 22, 1, 2, 246, 247, 5, 44, 23, 2, 247, 254, 3, 2, 2,
	2, 248, 249, 12, 3, 2, 2, 249, 250, 5, 56, 29, 2, 250, 251, 5, 44, 23,
	2, 251, 253, 3, 2, 2, 2, 252, 248, 3, 2, 2, 2, 253, 256, 3, 2, 2, 2, 254,
	252, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 43, 3, 2, 2, 2, 256, 254, 3,
	2, 2, 2, 257, 258, 8, 23, 1, 2, 258, 259, 5, 46, 24, 2, 259, 266, 3, 2,
	2, 2, 260, 261, 12, 3, 2, 2, 261, 262, 5, 58, 30, 2, 262, 263, 5, 46, 24,
	2, 263, 265, 3, 2, 2, 2, 264, 260, 3, 2, 2, 2, 265, 268, 3, 2, 2, 2, 266,
	264, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 45, 3, 2, 2, 2, 268, 266, 3,
	2, 2, 2, 269, 270, 8, 24, 1, 2, 270, 271, 5, 48, 25, 2, 271, 278, 3, 2,
	2, 2, 272, 273, 12, 3, 2, 2, 273, 274, 5, 60, 31, 2, 274, 275, 5, 48, 25,
	2, 275, 277, 3, 2, 2, 2, 276, 272, 3, 2, 2, 2, 277, 280, 3, 2, 2, 2, 278,
	276, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 47, 3, 2, 2, 2, 280, 278, 3,
	2, 2, 2, 281, 282, 8, 25, 1, 2, 282, 283, 5, 50, 26, 2, 283, 290, 3, 2,
	2, 2, 284, 285, 12, 3, 2, 2, 285, 286, 5, 62, 32, 2, 286, 287, 5, 50, 26,
	2, 287, 289, 3, 2, 2, 2, 288, 284, 3, 2, 2, 2, 289, 292, 3, 2, 2, 2, 290,
	288, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 49, 3, 2, 2, 2, 292, 290, 3,
	2, 2, 2, 293, 302, 5, 52, 27, 2, 294, 296, 5, 64, 33, 2, 295, 294, 3, 2,
	2, 2, 296, 297, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2,
	298, 299, 3, 2, 2, 2, 299, 300, 5, 50, 26, 2, 300, 302, 3, 2, 2, 2, 301,
	293, 3, 2, 2, 2, 301, 295, 3, 2, 2, 2, 302, 51, 3, 2, 2, 2, 303, 307, 5,
	66, 34, 2, 304, 306, 5, 54, 28, 2, 305, 304, 3, 2, 2, 2, 306, 309, 3, 2,
	2, 2, 307, 305, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 53, 3, 2, 2, 2,
	309, 307, 3, 2, 2, 2, 310, 313, 5, 68, 35, 2, 311, 313, 5, 74, 38, 2, 312,
	310, 3, 2, 2, 2, 312, 311, 3, 2, 2, 2, 313, 55, 3, 2, 2, 2, 314, 315, 9,
	3, 2, 2, 315, 57, 3, 2, 2, 2, 316, 317, 9, 4, 2, 2, 317, 59, 3, 2, 2, 2,
	318, 319, 9, 5, 2, 2, 319, 61, 3, 2, 2, 2, 320, 321, 9, 6, 2, 2, 321, 63,
	3, 2, 2, 2, 322, 323, 9, 7, 2, 2, 323, 65, 3, 2, 2, 2, 324, 342, 7, 40,
	2, 2, 325, 342, 5, 76, 39, 2, 326, 327, 7, 30, 2, 2, 327, 329, 7, 28, 2,
	2, 328, 330, 5, 8, 5, 2, 329, 328, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330,
	331, 3, 2, 2, 2, 331, 334, 7, 29, 2, 2, 332, 333, 7, 3, 2, 2, 333, 335,
	5, 12, 7, 2, 334, 332, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 336, 3, 2,
	2, 2, 336, 342, 5, 18, 10, 2, 337, 338, 7, 28, 2, 2, 338, 339, 5, 34, 18,
	2, 339, 340, 7, 29, 2, 2, 340, 342, 3, 2, 2, 2, 341, 324, 3, 2, 2, 2, 341,
	325, 3, 2, 2, 2, 341, 326, 3, 2, 2, 2, 341, 337, 3, 2, 2, 2, 342, 67, 3,
	2, 2, 2, 343, 346, 5, 70, 36, 2, 344, 346, 5, 72, 37, 2, 345, 343, 3, 2,
	2, 2, 345, 344, 3, 2, 2, 2, 346, 69, 3, 2, 2, 2, 347, 348, 7, 14, 2, 2,
	348, 349, 7, 40, 2, 2, 349, 71, 3, 2, 2, 2, 350, 351, 7, 6, 2, 2, 351,
	352, 5, 34, 18, 2, 352, 353, 7, 7, 2, 2, 353, 73, 3, 2, 2, 2, 354, 363,
	7, 28, 2, 2, 355, 360, 5, 34, 18, 2, 356, 357, 7, 4, 2, 2, 357, 359, 5,
	34, 18, 2, 358, 356, 3, 2, 2, 2, 359, 362, 3, 2, 2, 2, 360, 358, 3, 2,
	2, 2, 360, 361, 3, 2, 2, 2, 361, 364, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2,
	363, 355, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 365, 3, 2, 2, 2, 365,
	366, 7, 29, 2, 2, 366, 75, 3, 2, 2, 2, 367, 371, 5, 80, 41, 2, 368, 371,
	5, 78, 40, 2, 369, 371, 5, 82, 42, 2, 370, 367, 3, 2, 2, 2, 370, 368, 3,
	2, 2, 2, 370, 369, 3, 2, 2, 2, 371, 77, 3, 2, 2, 2, 372, 373, 9, 8, 2,
	2, 373, 79, 3, 2, 2, 2, 374, 379, 7, 41, 2, 2, 375, 379, 7, 42, 2, 2, 376,
	379, 7, 43, 2, 2, 377, 379, 7, 44, 2, 2, 378, 374, 3, 2, 2, 2, 378, 375,
	3, 2, 2, 2, 378, 376, 3, 2, 2, 2, 378, 377, 3, 2, 2, 2, 379, 81, 3, 2,
	2, 2, 380, 389, 7, 6, 2, 2, 381, 386, 5, 34, 18, 2, 382, 383, 7, 4, 2,
	2, 383, 385, 5, 34, 18, 2, 384, 382, 3, 2, 2, 2, 385, 388, 3, 2, 2, 2,
	386, 384, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 390, 3, 2, 2, 2, 388,
	386, 3, 2, 2, 2, 389, 381, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 391,
	3, 2, 2, 2, 391, 392, 7, 7, 2, 2, 392, 83, 3, 2, 2, 2, 393, 398, 7, 15,
	2, 2, 394, 398, 7, 2, 2, 3, 395, 398, 6, 43, 8, 2, 396, 398, 6, 43, 9,
	2, 397, 393, 3, 2, 2, 2, 397, 394, 3, 2, 2, 2, 397, 395, 3, 2, 2, 2, 397,
	396, 3, 2, 2, 2, 398, 85, 3, 2, 2, 2, 42, 89, 96, 99, 105, 110, 119, 130,
	135, 146, 149, 154, 165, 174, 178, 186, 188, 198, 207, 221, 231, 242, 254,
	266, 278, 290, 297, 301, 307, 312, 329, 334, 341, 345, 360, 363, 370, 378,
	386, 389, 397,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'->'", "','", "':'", "'['", "']'", "'{'", "'}'", "'='", "'?'", "'||'",
	"'&&'", "'.'", "';'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'+'",
	"'-'", "'*'", "'/'", "'%'", "'!'", "'('", "')'", "'fun'", "'pub'", "'return'",
	"'const'", "'var'", "'if'", "'else'", "'while'", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "Equal", "Unequal",
	"Less", "Greater", "LessEqual", "GreaterEqual", "Plus", "Minus", "Mul",
	"Div", "Mod", "Negate", "OpenParen", "CloseParen", "Fun", "Pub", "Return",
	"Const", "Var", "If", "Else", "While", "True", "False", "Identifier", "DecimalLiteral",
	"BinaryLiteral", "OctalLiteral", "HexadecimalLiteral", "WS", "Terminator",
	"BlockComment", "LineComment",
}

var ruleNames = []string{
	"program", "declaration", "functionDeclaration", "parameterList", "parameter",
	"typeName", "typeDimension", "baseType", "block", "statements", "statement",
	"returnStatement", "ifStatement", "whileStatement", "variableDeclaration",
	"assignment", "expression", "conditionalExpression", "orExpression", "andExpression",
	"equalityExpression", "relationalExpression", "additiveExpression", "multiplicativeExpression",
	"unaryExpression", "primaryExpression", "primaryExpressionSuffix", "equalityOp",
	"relationalOp", "additiveOp", "multiplicativeOp", "unaryOp", "primaryExpressionStart",
	"expressionAccess", "memberAccess", "bracketExpression", "invocation",
	"literal", "booleanLiteral", "integerLiteral", "arrayLiteral", "eos",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type StrictusParser struct {
	*antlr.BaseParser
}

func NewStrictusParser(input antlr.TokenStream) *StrictusParser {
	this := new(StrictusParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Strictus.g4"

	return this
}

// Returns true if on the current index of the parser's
// token stream a token exists on the Hidden channel which
// either is a line terminator, or is a multi line comment that
// contains a line terminator.
func (p *StrictusParser) lineTerminatorAhead() bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	if ahead.GetChannel() != antlr.LexerHidden {
		// We're only interested in tokens on the HIDDEN channel.
		return true
	}

	if ahead.GetTokenType() == StrictusParserTerminator {
		// There is definitely a line terminator ahead.
		return true
	}

	if ahead.GetTokenType() == StrictusParserWS {
		// Get the token ahead of the current whitespaces.
		possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - 2
		ahead = p.GetTokenStream().Get(possibleIndexEosToken)
	}

	// Get the token's text and type.
	text := ahead.GetText()
	_type := ahead.GetTokenType()

	// Check if the token is, or contains a line terminator.
	return (_type == StrictusParserBlockComment && (strings.Contains(text, "\r") || strings.Contains(text, "\n"))) ||
		(_type == StrictusParserTerminator)
}

// StrictusParser tokens.
const (
	StrictusParserEOF                = antlr.TokenEOF
	StrictusParserT__0               = 1
	StrictusParserT__1               = 2
	StrictusParserT__2               = 3
	StrictusParserT__3               = 4
	StrictusParserT__4               = 5
	StrictusParserT__5               = 6
	StrictusParserT__6               = 7
	StrictusParserT__7               = 8
	StrictusParserT__8               = 9
	StrictusParserT__9               = 10
	StrictusParserT__10              = 11
	StrictusParserT__11              = 12
	StrictusParserT__12              = 13
	StrictusParserEqual              = 14
	StrictusParserUnequal            = 15
	StrictusParserLess               = 16
	StrictusParserGreater            = 17
	StrictusParserLessEqual          = 18
	StrictusParserGreaterEqual       = 19
	StrictusParserPlus               = 20
	StrictusParserMinus              = 21
	StrictusParserMul                = 22
	StrictusParserDiv                = 23
	StrictusParserMod                = 24
	StrictusParserNegate             = 25
	StrictusParserOpenParen          = 26
	StrictusParserCloseParen         = 27
	StrictusParserFun                = 28
	StrictusParserPub                = 29
	StrictusParserReturn             = 30
	StrictusParserConst              = 31
	StrictusParserVar                = 32
	StrictusParserIf                 = 33
	StrictusParserElse               = 34
	StrictusParserWhile              = 35
	StrictusParserTrue               = 36
	StrictusParserFalse              = 37
	StrictusParserIdentifier         = 38
	StrictusParserDecimalLiteral     = 39
	StrictusParserBinaryLiteral      = 40
	StrictusParserOctalLiteral       = 41
	StrictusParserHexadecimalLiteral = 42
	StrictusParserWS                 = 43
	StrictusParserTerminator         = 44
	StrictusParserBlockComment       = 45
	StrictusParserLineComment        = 46
)

// StrictusParser rules.
const (
	StrictusParserRULE_program                  = 0
	StrictusParserRULE_declaration              = 1
	StrictusParserRULE_functionDeclaration      = 2
	StrictusParserRULE_parameterList            = 3
	StrictusParserRULE_parameter                = 4
	StrictusParserRULE_typeName                 = 5
	StrictusParserRULE_typeDimension            = 6
	StrictusParserRULE_baseType                 = 7
	StrictusParserRULE_block                    = 8
	StrictusParserRULE_statements               = 9
	StrictusParserRULE_statement                = 10
	StrictusParserRULE_returnStatement          = 11
	StrictusParserRULE_ifStatement              = 12
	StrictusParserRULE_whileStatement           = 13
	StrictusParserRULE_variableDeclaration      = 14
	StrictusParserRULE_assignment               = 15
	StrictusParserRULE_expression               = 16
	StrictusParserRULE_conditionalExpression    = 17
	StrictusParserRULE_orExpression             = 18
	StrictusParserRULE_andExpression            = 19
	StrictusParserRULE_equalityExpression       = 20
	StrictusParserRULE_relationalExpression     = 21
	StrictusParserRULE_additiveExpression       = 22
	StrictusParserRULE_multiplicativeExpression = 23
	StrictusParserRULE_unaryExpression          = 24
	StrictusParserRULE_primaryExpression        = 25
	StrictusParserRULE_primaryExpressionSuffix  = 26
	StrictusParserRULE_equalityOp               = 27
	StrictusParserRULE_relationalOp             = 28
	StrictusParserRULE_additiveOp               = 29
	StrictusParserRULE_multiplicativeOp         = 30
	StrictusParserRULE_unaryOp                  = 31
	StrictusParserRULE_primaryExpressionStart   = 32
	StrictusParserRULE_expressionAccess         = 33
	StrictusParserRULE_memberAccess             = 34
	StrictusParserRULE_bracketExpression        = 35
	StrictusParserRULE_invocation               = 36
	StrictusParserRULE_literal                  = 37
	StrictusParserRULE_booleanLiteral           = 38
	StrictusParserRULE_integerLiteral           = 39
	StrictusParserRULE_arrayLiteral             = 40
	StrictusParserRULE_eos                      = 41
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(StrictusParserEOF, 0)
}

func (s *ProgramContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *ProgramContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StrictusParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-28)&-(0x1f+1)) == 0 && ((1<<uint((_la-28)))&((1<<(StrictusParserFun-28))|(1<<(StrictusParserPub-28))|(1<<(StrictusParserConst-28))|(1<<(StrictusParserVar-28)))) != 0 {
		{
			p.SetState(84)
			p.Declaration()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(StrictusParserEOF)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *DeclarationContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StrictusParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(94)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserFun, StrictusParserPub:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.FunctionDeclaration()
		}

	case StrictusParserConst, StrictusParserVar:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.VariableDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeNameContext

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeNameContext)

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	returnType ITypeNameContext
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) GetReturnType() ITypeNameContext { return s.returnType }

func (s *FunctionDeclarationContext) SetReturnType(v ITypeNameContext) { s.returnType = v }

func (s *FunctionDeclarationContext) Fun() antlr.TerminalNode {
	return s.GetToken(StrictusParserFun, 0)
}

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, 0)
}

func (s *FunctionDeclarationContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, 0)
}

func (s *FunctionDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclarationContext) Pub() antlr.TerminalNode {
	return s.GetToken(StrictusParserPub, 0)
}

func (s *FunctionDeclarationContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionDeclarationContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, StrictusParserRULE_functionDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserPub {
		{
			p.SetState(96)
			p.Match(StrictusParserPub)
		}

	}
	{
		p.SetState(99)
		p.Match(StrictusParserFun)
	}
	{
		p.SetState(100)
		p.Match(StrictusParserIdentifier)
	}
	{
		p.SetState(101)
		p.Match(StrictusParserOpenParen)
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserIdentifier {
		{
			p.SetState(102)
			p.ParameterList()
		}

	}
	{
		p.SetState(105)
		p.Match(StrictusParserCloseParen)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserT__0 {
		{
			p.SetState(106)
			p.Match(StrictusParserT__0)
		}
		{
			p.SetState(107)

			var _x = p.TypeName()

			localctx.(*FunctionDeclarationContext).returnType = _x
		}

	}
	{
		p.SetState(110)
		p.Block()
	}

	return localctx
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_parameterList
	return p
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, StrictusParserRULE_parameterList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Parameter()
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == StrictusParserT__1 {
		{
			p.SetState(113)
			p.Match(StrictusParserT__1)
		}
		{
			p.SetState(114)
			p.Parameter()
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *ParameterContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, StrictusParserRULE_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(StrictusParserIdentifier)
	}
	{
		p.SetState(121)
		p.Match(StrictusParserT__2)
	}
	{
		p.SetState(122)
		p.TypeName()
	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) BaseType() IBaseTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *TypeNameContext) AllTypeDimension() []ITypeDimensionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeDimensionContext)(nil)).Elem())
	var tst = make([]ITypeDimensionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeDimensionContext)
		}
	}

	return tst
}

func (s *TypeNameContext) TypeDimension(i int) ITypeDimensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDimensionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeDimensionContext)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, StrictusParserRULE_typeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.BaseType()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(125)
				p.TypeDimension()
			}

		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeDimensionContext is an interface to support dynamic dispatch.
type ITypeDimensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDimensionContext differentiates from other interfaces.
	IsTypeDimensionContext()
}

type TypeDimensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDimensionContext() *TypeDimensionContext {
	var p = new(TypeDimensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_typeDimension
	return p
}

func (*TypeDimensionContext) IsTypeDimensionContext() {}

func NewTypeDimensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDimensionContext {
	var p = new(TypeDimensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_typeDimension

	return p
}

func (s *TypeDimensionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDimensionContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserDecimalLiteral, 0)
}

func (s *TypeDimensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDimensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDimensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterTypeDimension(s)
	}
}

func (s *TypeDimensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitTypeDimension(s)
	}
}

func (s *TypeDimensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitTypeDimension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) TypeDimension() (localctx ITypeDimensionContext) {
	localctx = NewTypeDimensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, StrictusParserRULE_typeDimension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(StrictusParserT__3)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserDecimalLiteral {
		{
			p.SetState(132)
			p.Match(StrictusParserDecimalLiteral)
		}

	}
	{
		p.SetState(135)
		p.Match(StrictusParserT__4)
	}

	return localctx
}

// IBaseTypeContext is an interface to support dynamic dispatch.
type IBaseTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_typeName returns the _typeName rule contexts.
	Get_typeName() ITypeNameContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeNameContext

	// Set_typeName sets the _typeName rule contexts.
	Set_typeName(ITypeNameContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeNameContext)

	// GetParameterTypes returns the parameterTypes rule context list.
	GetParameterTypes() []ITypeNameContext

	// SetParameterTypes sets the parameterTypes rule context list.
	SetParameterTypes([]ITypeNameContext)

	// IsBaseTypeContext differentiates from other interfaces.
	IsBaseTypeContext()
}

type BaseTypeContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_typeName      ITypeNameContext
	parameterTypes []ITypeNameContext
	returnType     ITypeNameContext
}

func NewEmptyBaseTypeContext() *BaseTypeContext {
	var p = new(BaseTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_baseType
	return p
}

func (*BaseTypeContext) IsBaseTypeContext() {}

func NewBaseTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseTypeContext {
	var p = new(BaseTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_baseType

	return p
}

func (s *BaseTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseTypeContext) Get_typeName() ITypeNameContext { return s._typeName }

func (s *BaseTypeContext) GetReturnType() ITypeNameContext { return s.returnType }

func (s *BaseTypeContext) Set_typeName(v ITypeNameContext) { s._typeName = v }

func (s *BaseTypeContext) SetReturnType(v ITypeNameContext) { s.returnType = v }

func (s *BaseTypeContext) GetParameterTypes() []ITypeNameContext { return s.parameterTypes }

func (s *BaseTypeContext) SetParameterTypes(v []ITypeNameContext) { s.parameterTypes = v }

func (s *BaseTypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *BaseTypeContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, 0)
}

func (s *BaseTypeContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, 0)
}

func (s *BaseTypeContext) AllTypeName() []ITypeNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeNameContext)(nil)).Elem())
	var tst = make([]ITypeNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeNameContext)
		}
	}

	return tst
}

func (s *BaseTypeContext) TypeName(i int) ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (s *BaseTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBaseType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) BaseType() (localctx IBaseTypeContext) {
	localctx = NewBaseTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, StrictusParserRULE_baseType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Match(StrictusParserIdentifier)
		}

	case StrictusParserOpenParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.Match(StrictusParserOpenParen)
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StrictusParserOpenParen || _la == StrictusParserIdentifier {
			{
				p.SetState(139)

				var _x = p.TypeName()

				localctx.(*BaseTypeContext)._typeName = _x
			}
			localctx.(*BaseTypeContext).parameterTypes = append(localctx.(*BaseTypeContext).parameterTypes, localctx.(*BaseTypeContext)._typeName)
			p.SetState(144)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StrictusParserT__1 {
				{
					p.SetState(140)
					p.Match(StrictusParserT__1)
				}
				{
					p.SetState(141)

					var _x = p.TypeName()

					localctx.(*BaseTypeContext)._typeName = _x
				}
				localctx.(*BaseTypeContext).parameterTypes = append(localctx.(*BaseTypeContext).parameterTypes, localctx.(*BaseTypeContext)._typeName)

				p.SetState(146)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(149)
			p.Match(StrictusParserCloseParen)
		}
		{
			p.SetState(150)
			p.Match(StrictusParserT__0)
		}
		{
			p.SetState(151)

			var _x = p.TypeName()

			localctx.(*BaseTypeContext).returnType = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, StrictusParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(StrictusParserT__5)
	}
	{
		p.SetState(155)
		p.Statements()
	}
	{
		p.SetState(156)
		p.Match(StrictusParserT__6)
	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *StatementsContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, StrictusParserRULE_statements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__3)|(1<<StrictusParserMinus)|(1<<StrictusParserNegate)|(1<<StrictusParserOpenParen)|(1<<StrictusParserFun)|(1<<StrictusParserPub)|(1<<StrictusParserReturn)|(1<<StrictusParserConst))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(StrictusParserVar-32))|(1<<(StrictusParserIf-32))|(1<<(StrictusParserWhile-32))|(1<<(StrictusParserTrue-32))|(1<<(StrictusParserFalse-32))|(1<<(StrictusParserIdentifier-32))|(1<<(StrictusParserDecimalLiteral-32))|(1<<(StrictusParserBinaryLiteral-32))|(1<<(StrictusParserOctalLiteral-32))|(1<<(StrictusParserHexadecimalLiteral-32)))) != 0) {
		{
			p.SetState(158)
			p.Statement()
		}
		{
			p.SetState(159)
			p.Eos()
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, StrictusParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.ReturnStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.IfStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
			p.WhileStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(169)
			p.Declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(170)
			p.Assignment()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(171)
			p.Expression()
		}

	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(StrictusParserReturn, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, StrictusParserRULE_returnStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(StrictusParserReturn)
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(175)
			p.Expression()
		}

	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTest returns the test rule contexts.
	GetTest() IExpressionContext

	// GetThen returns the then rule contexts.
	GetThen() IBlockContext

	// GetAlt returns the alt rule contexts.
	GetAlt() IBlockContext

	// SetTest sets the test rule contexts.
	SetTest(IExpressionContext)

	// SetThen sets the then rule contexts.
	SetThen(IBlockContext)

	// SetAlt sets the alt rule contexts.
	SetAlt(IBlockContext)

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	test   IExpressionContext
	then   IBlockContext
	alt    IBlockContext
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) GetTest() IExpressionContext { return s.test }

func (s *IfStatementContext) GetThen() IBlockContext { return s.then }

func (s *IfStatementContext) GetAlt() IBlockContext { return s.alt }

func (s *IfStatementContext) SetTest(v IExpressionContext) { s.test = v }

func (s *IfStatementContext) SetThen(v IBlockContext) { s.then = v }

func (s *IfStatementContext) SetAlt(v IBlockContext) { s.alt = v }

func (s *IfStatementContext) If() antlr.TerminalNode {
	return s.GetToken(StrictusParserIf, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(StrictusParserElse, 0)
}

func (s *IfStatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, StrictusParserRULE_ifStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(StrictusParserIf)
	}
	{
		p.SetState(179)

		var _x = p.Expression()

		localctx.(*IfStatementContext).test = _x
	}
	{
		p.SetState(180)

		var _x = p.Block()

		localctx.(*IfStatementContext).then = _x
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(181)
			p.Match(StrictusParserElse)
		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case StrictusParserIf:
			{
				p.SetState(182)
				p.IfStatement()
			}

		case StrictusParserT__5:
			{
				p.SetState(183)

				var _x = p.Block()

				localctx.(*IfStatementContext).alt = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_whileStatement
	return p
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(StrictusParserWhile, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, StrictusParserRULE_whileStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(StrictusParserWhile)
	}
	{
		p.SetState(189)
		p.Expression()
	}
	{
		p.SetState(190)
		p.Block()
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *VariableDeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationContext) Const() antlr.TerminalNode {
	return s.GetToken(StrictusParserConst, 0)
}

func (s *VariableDeclarationContext) Var() antlr.TerminalNode {
	return s.GetToken(StrictusParserVar, 0)
}

func (s *VariableDeclarationContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, StrictusParserRULE_variableDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserConst || _la == StrictusParserVar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(193)
		p.Match(StrictusParserIdentifier)
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserT__2 {
		{
			p.SetState(194)
			p.Match(StrictusParserT__2)
		}
		{
			p.SetState(195)
			p.TypeName()
		}

	}
	{
		p.SetState(198)
		p.Match(StrictusParserT__7)
	}
	{
		p.SetState(199)
		p.Expression()
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) AllExpressionAccess() []IExpressionAccessContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAccessContext)(nil)).Elem())
	var tst = make([]IExpressionAccessContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAccessContext)
		}
	}

	return tst
}

func (s *AssignmentContext) ExpressionAccess(i int) IExpressionAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAccessContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAccessContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, StrictusParserRULE_assignment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(StrictusParserIdentifier)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == StrictusParserT__3 || _la == StrictusParserT__11 {
		{
			p.SetState(202)
			p.ExpressionAccess()
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(208)
		p.Match(StrictusParserT__7)
	}
	{
		p.SetState(209)
		p.Expression()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ConditionalExpression() IConditionalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, StrictusParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.ConditionalExpression()
	}

	return localctx
}

// IConditionalExpressionContext is an interface to support dynamic dispatch.
type IConditionalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetThen returns the then rule contexts.
	GetThen() IExpressionContext

	// GetAlt returns the alt rule contexts.
	GetAlt() IExpressionContext

	// SetThen sets the then rule contexts.
	SetThen(IExpressionContext)

	// SetAlt sets the alt rule contexts.
	SetAlt(IExpressionContext)

	// IsConditionalExpressionContext differentiates from other interfaces.
	IsConditionalExpressionContext()
}

type ConditionalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	then   IExpressionContext
	alt    IExpressionContext
}

func NewEmptyConditionalExpressionContext() *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_conditionalExpression
	return p
}

func (*ConditionalExpressionContext) IsConditionalExpressionContext() {}

func NewConditionalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_conditionalExpression

	return p
}

func (s *ConditionalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalExpressionContext) GetThen() IExpressionContext { return s.then }

func (s *ConditionalExpressionContext) GetAlt() IExpressionContext { return s.alt }

func (s *ConditionalExpressionContext) SetThen(v IExpressionContext) { s.then = v }

func (s *ConditionalExpressionContext) SetAlt(v IExpressionContext) { s.alt = v }

func (s *ConditionalExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *ConditionalExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionalExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitConditionalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ConditionalExpression() (localctx IConditionalExpressionContext) {
	localctx = NewConditionalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, StrictusParserRULE_conditionalExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.orExpression(0)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(214)
			p.Match(StrictusParserT__8)
		}
		{
			p.SetState(215)

			var _x = p.Expression()

			localctx.(*ConditionalExpressionContext).then = _x
		}
		{
			p.SetState(216)
			p.Match(StrictusParserT__2)
		}
		{
			p.SetState(217)

			var _x = p.Expression()

			localctx.(*ConditionalExpressionContext).alt = _x
		}

	}

	return localctx
}

// IOrExpressionContext is an interface to support dynamic dispatch.
type IOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrExpressionContext differentiates from other interfaces.
	IsOrExpressionContext()
}

type OrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExpressionContext() *OrExpressionContext {
	var p = new(OrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_orExpression
	return p
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_orExpression

	return p
}

func (s *OrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *OrExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (s *OrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) OrExpression() (localctx IOrExpressionContext) {
	return p.orExpression(0)
}

func (p *StrictusParser) orExpression(_p int) (localctx IOrExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IOrExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, StrictusParserRULE_orExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.andExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOrExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_orExpression)
			p.SetState(224)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(225)
				p.Match(StrictusParserT__9)
			}
			{
				p.SetState(226)
				p.andExpression(0)
			}

		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_andExpression
	return p
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *AndExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) AndExpression() (localctx IAndExpressionContext) {
	return p.andExpression(0)
}

func (p *StrictusParser) andExpression(_p int) (localctx IAndExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAndExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, StrictusParserRULE_andExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.equalityExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_andExpression)
			p.SetState(235)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(236)
				p.Match(StrictusParserT__10)
			}
			{
				p.SetState(237)
				p.equalityExpression(0)
			}

		}
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_equalityExpression
	return p
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *EqualityExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *EqualityExpressionContext) EqualityOp() IEqualityOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityOpContext)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	return p.equalityExpression(0)
}

func (p *StrictusParser) equalityExpression(_p int) (localctx IEqualityExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEqualityExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, StrictusParserRULE_equalityExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.relationalExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewEqualityExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_equalityExpression)
			p.SetState(246)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(247)
				p.EqualityOp()
			}
			{
				p.SetState(248)
				p.relationalExpression(0)
			}

		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}

	return localctx
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_relationalExpression
	return p
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *RelationalExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *RelationalExpressionContext) RelationalOp() IRelationalOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalOpContext)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	return p.relationalExpression(0)
}

func (p *StrictusParser) relationalExpression(_p int) (localctx IRelationalExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelationalExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, StrictusParserRULE_relationalExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.additiveExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_relationalExpression)
			p.SetState(258)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(259)
				p.RelationalOp()
			}
			{
				p.SetState(260)
				p.additiveExpression(0)
			}

		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_additiveExpression
	return p
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) MultiplicativeExpression() IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *AdditiveExpressionContext) AdditiveOp() IAdditiveOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveOpContext)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	return p.additiveExpression(0)
}

func (p *StrictusParser) additiveExpression(_p int) (localctx IAdditiveExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAdditiveExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, StrictusParserRULE_additiveExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.multiplicativeExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAdditiveExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_additiveExpression)
			p.SetState(270)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(271)
				p.AdditiveOp()
			}
			{
				p.SetState(272)
				p.multiplicativeExpression(0)
			}

		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_multiplicativeExpression
	return p
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicativeExpressionContext) MultiplicativeExpression() IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *MultiplicativeExpressionContext) MultiplicativeOp() IMultiplicativeOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeOpContext)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	return p.multiplicativeExpression(0)
}

func (p *StrictusParser) multiplicativeExpression(_p int) (localctx IMultiplicativeExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultiplicativeExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, StrictusParserRULE_multiplicativeExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.UnaryExpression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMultiplicativeExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_multiplicativeExpression)
			p.SetState(282)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(283)
				p.MultiplicativeOp()
			}
			{
				p.SetState(284)
				p.UnaryExpression()
			}

		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_unaryExpression
	return p
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *UnaryExpressionContext) AllUnaryOp() []IUnaryOpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnaryOpContext)(nil)).Elem())
	var tst = make([]IUnaryOpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnaryOpContext)
		}
	}

	return tst
}

func (s *UnaryExpressionContext) UnaryOp(i int) IUnaryOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, StrictusParserRULE_unaryExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(299)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserT__3, StrictusParserOpenParen, StrictusParserFun, StrictusParserTrue, StrictusParserFalse, StrictusParserIdentifier, StrictusParserDecimalLiteral, StrictusParserBinaryLiteral, StrictusParserOctalLiteral, StrictusParserHexadecimalLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(291)
			p.PrimaryExpression()
		}

	case StrictusParserMinus, StrictusParserNegate:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(292)
					p.UnaryOp()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(295)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
		}
		{
			p.SetState(297)
			p.UnaryExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_primaryExpression
	return p
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) PrimaryExpressionStart() IPrimaryExpressionStartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionStartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionStartContext)
}

func (s *PrimaryExpressionContext) AllPrimaryExpressionSuffix() []IPrimaryExpressionSuffixContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimaryExpressionSuffixContext)(nil)).Elem())
	var tst = make([]IPrimaryExpressionSuffixContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimaryExpressionSuffixContext)
		}
	}

	return tst
}

func (s *PrimaryExpressionContext) PrimaryExpressionSuffix(i int) IPrimaryExpressionSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionSuffixContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionSuffixContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, StrictusParserRULE_primaryExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.PrimaryExpressionStart()
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(302)
				p.PrimaryExpressionSuffix()
			}

		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExpressionSuffixContext is an interface to support dynamic dispatch.
type IPrimaryExpressionSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionSuffixContext differentiates from other interfaces.
	IsPrimaryExpressionSuffixContext()
}

type PrimaryExpressionSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionSuffixContext() *PrimaryExpressionSuffixContext {
	var p = new(PrimaryExpressionSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_primaryExpressionSuffix
	return p
}

func (*PrimaryExpressionSuffixContext) IsPrimaryExpressionSuffixContext() {}

func NewPrimaryExpressionSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionSuffixContext {
	var p = new(PrimaryExpressionSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_primaryExpressionSuffix

	return p
}

func (s *PrimaryExpressionSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionSuffixContext) ExpressionAccess() IExpressionAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAccessContext)
}

func (s *PrimaryExpressionSuffixContext) Invocation() IInvocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *PrimaryExpressionSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterPrimaryExpressionSuffix(s)
	}
}

func (s *PrimaryExpressionSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitPrimaryExpressionSuffix(s)
	}
}

func (s *PrimaryExpressionSuffixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitPrimaryExpressionSuffix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PrimaryExpressionSuffix() (localctx IPrimaryExpressionSuffixContext) {
	localctx = NewPrimaryExpressionSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, StrictusParserRULE_primaryExpressionSuffix)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(310)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserT__3, StrictusParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(308)
			p.ExpressionAccess()
		}

	case StrictusParserOpenParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(309)
			p.Invocation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEqualityOpContext is an interface to support dynamic dispatch.
type IEqualityOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualityOpContext differentiates from other interfaces.
	IsEqualityOpContext()
}

type EqualityOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityOpContext() *EqualityOpContext {
	var p = new(EqualityOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_equalityOp
	return p
}

func (*EqualityOpContext) IsEqualityOpContext() {}

func NewEqualityOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityOpContext {
	var p = new(EqualityOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_equalityOp

	return p
}

func (s *EqualityOpContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityOpContext) Equal() antlr.TerminalNode {
	return s.GetToken(StrictusParserEqual, 0)
}

func (s *EqualityOpContext) Unequal() antlr.TerminalNode {
	return s.GetToken(StrictusParserUnequal, 0)
}

func (s *EqualityOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterEqualityOp(s)
	}
}

func (s *EqualityOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitEqualityOp(s)
	}
}

func (s *EqualityOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitEqualityOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) EqualityOp() (localctx IEqualityOpContext) {
	localctx = NewEqualityOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, StrictusParserRULE_equalityOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserEqual || _la == StrictusParserUnequal) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRelationalOpContext is an interface to support dynamic dispatch.
type IRelationalOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalOpContext differentiates from other interfaces.
	IsRelationalOpContext()
}

type RelationalOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalOpContext() *RelationalOpContext {
	var p = new(RelationalOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_relationalOp
	return p
}

func (*RelationalOpContext) IsRelationalOpContext() {}

func NewRelationalOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalOpContext {
	var p = new(RelationalOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_relationalOp

	return p
}

func (s *RelationalOpContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalOpContext) Less() antlr.TerminalNode {
	return s.GetToken(StrictusParserLess, 0)
}

func (s *RelationalOpContext) Greater() antlr.TerminalNode {
	return s.GetToken(StrictusParserGreater, 0)
}

func (s *RelationalOpContext) LessEqual() antlr.TerminalNode {
	return s.GetToken(StrictusParserLessEqual, 0)
}

func (s *RelationalOpContext) GreaterEqual() antlr.TerminalNode {
	return s.GetToken(StrictusParserGreaterEqual, 0)
}

func (s *RelationalOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterRelationalOp(s)
	}
}

func (s *RelationalOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitRelationalOp(s)
	}
}

func (s *RelationalOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitRelationalOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) RelationalOp() (localctx IRelationalOpContext) {
	localctx = NewRelationalOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, StrictusParserRULE_relationalOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserLess)|(1<<StrictusParserGreater)|(1<<StrictusParserLessEqual)|(1<<StrictusParserGreaterEqual))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAdditiveOpContext is an interface to support dynamic dispatch.
type IAdditiveOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveOpContext differentiates from other interfaces.
	IsAdditiveOpContext()
}

type AdditiveOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveOpContext() *AdditiveOpContext {
	var p = new(AdditiveOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_additiveOp
	return p
}

func (*AdditiveOpContext) IsAdditiveOpContext() {}

func NewAdditiveOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveOpContext {
	var p = new(AdditiveOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_additiveOp

	return p
}

func (s *AdditiveOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveOpContext) Plus() antlr.TerminalNode {
	return s.GetToken(StrictusParserPlus, 0)
}

func (s *AdditiveOpContext) Minus() antlr.TerminalNode {
	return s.GetToken(StrictusParserMinus, 0)
}

func (s *AdditiveOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAdditiveOp(s)
	}
}

func (s *AdditiveOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAdditiveOp(s)
	}
}

func (s *AdditiveOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAdditiveOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) AdditiveOp() (localctx IAdditiveOpContext) {
	localctx = NewAdditiveOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, StrictusParserRULE_additiveOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserPlus || _la == StrictusParserMinus) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMultiplicativeOpContext is an interface to support dynamic dispatch.
type IMultiplicativeOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeOpContext differentiates from other interfaces.
	IsMultiplicativeOpContext()
}

type MultiplicativeOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeOpContext() *MultiplicativeOpContext {
	var p = new(MultiplicativeOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_multiplicativeOp
	return p
}

func (*MultiplicativeOpContext) IsMultiplicativeOpContext() {}

func NewMultiplicativeOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeOpContext {
	var p = new(MultiplicativeOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_multiplicativeOp

	return p
}

func (s *MultiplicativeOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeOpContext) Mul() antlr.TerminalNode {
	return s.GetToken(StrictusParserMul, 0)
}

func (s *MultiplicativeOpContext) Div() antlr.TerminalNode {
	return s.GetToken(StrictusParserDiv, 0)
}

func (s *MultiplicativeOpContext) Mod() antlr.TerminalNode {
	return s.GetToken(StrictusParserMod, 0)
}

func (s *MultiplicativeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterMultiplicativeOp(s)
	}
}

func (s *MultiplicativeOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitMultiplicativeOp(s)
	}
}

func (s *MultiplicativeOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitMultiplicativeOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) MultiplicativeOp() (localctx IMultiplicativeOpContext) {
	localctx = NewMultiplicativeOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, StrictusParserRULE_multiplicativeOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserMul)|(1<<StrictusParserDiv)|(1<<StrictusParserMod))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnaryOpContext is an interface to support dynamic dispatch.
type IUnaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOpContext differentiates from other interfaces.
	IsUnaryOpContext()
}

type UnaryOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpContext() *UnaryOpContext {
	var p = new(UnaryOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_unaryOp
	return p
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOpContext) Minus() antlr.TerminalNode {
	return s.GetToken(StrictusParserMinus, 0)
}

func (s *UnaryOpContext) Negate() antlr.TerminalNode {
	return s.GetToken(StrictusParserNegate, 0)
}

func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterUnaryOp(s)
	}
}

func (s *UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitUnaryOp(s)
	}
}

func (s *UnaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitUnaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) UnaryOp() (localctx IUnaryOpContext) {
	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, StrictusParserRULE_unaryOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserMinus || _la == StrictusParserNegate) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPrimaryExpressionStartContext is an interface to support dynamic dispatch.
type IPrimaryExpressionStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionStartContext differentiates from other interfaces.
	IsPrimaryExpressionStartContext()
}

type PrimaryExpressionStartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionStartContext() *PrimaryExpressionStartContext {
	var p = new(PrimaryExpressionStartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_primaryExpressionStart
	return p
}

func (*PrimaryExpressionStartContext) IsPrimaryExpressionStartContext() {}

func NewPrimaryExpressionStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionStartContext {
	var p = new(PrimaryExpressionStartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_primaryExpressionStart

	return p
}

func (s *PrimaryExpressionStartContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionStartContext) CopyFrom(ctx *PrimaryExpressionStartContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PrimaryExpressionStartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionStartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionExpressionContext struct {
	*PrimaryExpressionStartContext
	returnType ITypeNameContext
}

func NewFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	p.PrimaryExpressionStartContext = NewEmptyPrimaryExpressionStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionStartContext))

	return p
}

func (s *FunctionExpressionContext) GetReturnType() ITypeNameContext { return s.returnType }

func (s *FunctionExpressionContext) SetReturnType(v ITypeNameContext) { s.returnType = v }

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) Fun() antlr.TerminalNode {
	return s.GetToken(StrictusParserFun, 0)
}

func (s *FunctionExpressionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, 0)
}

func (s *FunctionExpressionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, 0)
}

func (s *FunctionExpressionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionExpressionContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionExpressionContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFunctionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NestedExpressionContext struct {
	*PrimaryExpressionStartContext
}

func NewNestedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedExpressionContext {
	var p = new(NestedExpressionContext)

	p.PrimaryExpressionStartContext = NewEmptyPrimaryExpressionStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionStartContext))

	return p
}

func (s *NestedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedExpressionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, 0)
}

func (s *NestedExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NestedExpressionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, 0)
}

func (s *NestedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterNestedExpression(s)
	}
}

func (s *NestedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitNestedExpression(s)
	}
}

func (s *NestedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitNestedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExpressionContext struct {
	*PrimaryExpressionStartContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.PrimaryExpressionStartContext = NewEmptyPrimaryExpressionStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionStartContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	*PrimaryExpressionStartContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.PrimaryExpressionStartContext = NewEmptyPrimaryExpressionStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionStartContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PrimaryExpressionStart() (localctx IPrimaryExpressionStartContext) {
	localctx = NewPrimaryExpressionStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, StrictusParserRULE_primaryExpressionStart)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(339)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserIdentifier:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.Match(StrictusParserIdentifier)
		}

	case StrictusParserT__3, StrictusParserTrue, StrictusParserFalse, StrictusParserDecimalLiteral, StrictusParserBinaryLiteral, StrictusParserOctalLiteral, StrictusParserHexadecimalLiteral:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(323)
			p.Literal()
		}

	case StrictusParserFun:
		localctx = NewFunctionExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(324)
			p.Match(StrictusParserFun)
		}
		{
			p.SetState(325)
			p.Match(StrictusParserOpenParen)
		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StrictusParserIdentifier {
			{
				p.SetState(326)
				p.ParameterList()
			}

		}
		{
			p.SetState(329)
			p.Match(StrictusParserCloseParen)
		}
		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StrictusParserT__0 {
			{
				p.SetState(330)
				p.Match(StrictusParserT__0)
			}
			{
				p.SetState(331)

				var _x = p.TypeName()

				localctx.(*FunctionExpressionContext).returnType = _x
			}

		}
		{
			p.SetState(334)
			p.Block()
		}

	case StrictusParserOpenParen:
		localctx = NewNestedExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(335)
			p.Match(StrictusParserOpenParen)
		}
		{
			p.SetState(336)
			p.Expression()
		}
		{
			p.SetState(337)
			p.Match(StrictusParserCloseParen)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionAccessContext is an interface to support dynamic dispatch.
type IExpressionAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionAccessContext differentiates from other interfaces.
	IsExpressionAccessContext()
}

type ExpressionAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAccessContext() *ExpressionAccessContext {
	var p = new(ExpressionAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_expressionAccess
	return p
}

func (*ExpressionAccessContext) IsExpressionAccessContext() {}

func NewExpressionAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAccessContext {
	var p = new(ExpressionAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_expressionAccess

	return p
}

func (s *ExpressionAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAccessContext) MemberAccess() IMemberAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberAccessContext)
}

func (s *ExpressionAccessContext) BracketExpression() IBracketExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBracketExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBracketExpressionContext)
}

func (s *ExpressionAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterExpressionAccess(s)
	}
}

func (s *ExpressionAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitExpressionAccess(s)
	}
}

func (s *ExpressionAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitExpressionAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ExpressionAccess() (localctx IExpressionAccessContext) {
	localctx = NewExpressionAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, StrictusParserRULE_expressionAccess)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(343)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(341)
			p.MemberAccess()
		}

	case StrictusParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.BracketExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMemberAccessContext is an interface to support dynamic dispatch.
type IMemberAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberAccessContext differentiates from other interfaces.
	IsMemberAccessContext()
}

type MemberAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberAccessContext() *MemberAccessContext {
	var p = new(MemberAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_memberAccess
	return p
}

func (*MemberAccessContext) IsMemberAccessContext() {}

func NewMemberAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberAccessContext {
	var p = new(MemberAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_memberAccess

	return p
}

func (s *MemberAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberAccessContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *MemberAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterMemberAccess(s)
	}
}

func (s *MemberAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitMemberAccess(s)
	}
}

func (s *MemberAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitMemberAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) MemberAccess() (localctx IMemberAccessContext) {
	localctx = NewMemberAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, StrictusParserRULE_memberAccess)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(StrictusParserT__11)
	}
	{
		p.SetState(346)
		p.Match(StrictusParserIdentifier)
	}

	return localctx
}

// IBracketExpressionContext is an interface to support dynamic dispatch.
type IBracketExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBracketExpressionContext differentiates from other interfaces.
	IsBracketExpressionContext()
}

type BracketExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketExpressionContext() *BracketExpressionContext {
	var p = new(BracketExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_bracketExpression
	return p
}

func (*BracketExpressionContext) IsBracketExpressionContext() {}

func NewBracketExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracketExpressionContext {
	var p = new(BracketExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_bracketExpression

	return p
}

func (s *BracketExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BracketExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BracketExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BracketExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBracketExpression(s)
	}
}

func (s *BracketExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBracketExpression(s)
	}
}

func (s *BracketExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBracketExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) BracketExpression() (localctx IBracketExpressionContext) {
	localctx = NewBracketExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, StrictusParserRULE_bracketExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(348)
		p.Match(StrictusParserT__3)
	}
	{
		p.SetState(349)
		p.Expression()
	}
	{
		p.SetState(350)
		p.Match(StrictusParserT__4)
	}

	return localctx
}

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_invocation
	return p
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, 0)
}

func (s *InvocationContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, 0)
}

func (s *InvocationContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *InvocationContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterInvocation(s)
	}
}

func (s *InvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitInvocation(s)
	}
}

func (s *InvocationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitInvocation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Invocation() (localctx IInvocationContext) {
	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, StrictusParserRULE_invocation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(StrictusParserOpenParen)
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__3)|(1<<StrictusParserMinus)|(1<<StrictusParserNegate)|(1<<StrictusParserOpenParen)|(1<<StrictusParserFun))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(StrictusParserTrue-36))|(1<<(StrictusParserFalse-36))|(1<<(StrictusParserIdentifier-36))|(1<<(StrictusParserDecimalLiteral-36))|(1<<(StrictusParserBinaryLiteral-36))|(1<<(StrictusParserOctalLiteral-36))|(1<<(StrictusParserHexadecimalLiteral-36)))) != 0) {
		{
			p.SetState(353)
			p.Expression()
		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(354)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(355)
				p.Expression()
			}

			p.SetState(360)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(363)
		p.Match(StrictusParserCloseParen)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *LiteralContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *LiteralContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, StrictusParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(368)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserDecimalLiteral, StrictusParserBinaryLiteral, StrictusParserOctalLiteral, StrictusParserHexadecimalLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(365)
			p.IntegerLiteral()
		}

	case StrictusParserTrue, StrictusParserFalse:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(366)
			p.BooleanLiteral()
		}

	case StrictusParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(367)
			p.ArrayLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) True() antlr.TerminalNode {
	return s.GetToken(StrictusParserTrue, 0)
}

func (s *BooleanLiteralContext) False() antlr.TerminalNode {
	return s.GetToken(StrictusParserFalse, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, StrictusParserRULE_booleanLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserTrue || _la == StrictusParserFalse) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) CopyFrom(ctx *IntegerLiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryLiteralContext struct {
	*IntegerLiteralContext
}

func NewBinaryLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryLiteralContext {
	var p = new(BinaryLiteralContext)

	p.IntegerLiteralContext = NewEmptyIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IntegerLiteralContext))

	return p
}

func (s *BinaryLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLiteralContext) BinaryLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserBinaryLiteral, 0)
}

func (s *BinaryLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBinaryLiteral(s)
	}
}

func (s *BinaryLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBinaryLiteral(s)
	}
}

func (s *BinaryLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBinaryLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type OctalLiteralContext struct {
	*IntegerLiteralContext
}

func NewOctalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OctalLiteralContext {
	var p = new(OctalLiteralContext)

	p.IntegerLiteralContext = NewEmptyIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IntegerLiteralContext))

	return p
}

func (s *OctalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OctalLiteralContext) OctalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserOctalLiteral, 0)
}

func (s *OctalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterOctalLiteral(s)
	}
}

func (s *OctalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitOctalLiteral(s)
	}
}

func (s *OctalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitOctalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecimalLiteralContext struct {
	*IntegerLiteralContext
}

func NewDecimalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.IntegerLiteralContext = NewEmptyIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IntegerLiteralContext))

	return p
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserDecimalLiteral, 0)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitDecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type HexadecimalLiteralContext struct {
	*IntegerLiteralContext
}

func NewHexadecimalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HexadecimalLiteralContext {
	var p = new(HexadecimalLiteralContext)

	p.IntegerLiteralContext = NewEmptyIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IntegerLiteralContext))

	return p
}

func (s *HexadecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexadecimalLiteralContext) HexadecimalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserHexadecimalLiteral, 0)
}

func (s *HexadecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterHexadecimalLiteral(s)
	}
}

func (s *HexadecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitHexadecimalLiteral(s)
	}
}

func (s *HexadecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitHexadecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, StrictusParserRULE_integerLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(376)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserDecimalLiteral:
		localctx = NewDecimalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(372)
			p.Match(StrictusParserDecimalLiteral)
		}

	case StrictusParserBinaryLiteral:
		localctx = NewBinaryLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(373)
			p.Match(StrictusParserBinaryLiteral)
		}

	case StrictusParserOctalLiteral:
		localctx = NewOctalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(374)
			p.Match(StrictusParserOctalLiteral)
		}

	case StrictusParserHexadecimalLiteral:
		localctx = NewHexadecimalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(375)
			p.Match(StrictusParserHexadecimalLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, StrictusParserRULE_arrayLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(StrictusParserT__3)
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__3)|(1<<StrictusParserMinus)|(1<<StrictusParserNegate)|(1<<StrictusParserOpenParen)|(1<<StrictusParserFun))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(StrictusParserTrue-36))|(1<<(StrictusParserFalse-36))|(1<<(StrictusParserIdentifier-36))|(1<<(StrictusParserDecimalLiteral-36))|(1<<(StrictusParserBinaryLiteral-36))|(1<<(StrictusParserOctalLiteral-36))|(1<<(StrictusParserHexadecimalLiteral-36)))) != 0) {
		{
			p.SetState(379)
			p.Expression()
		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(380)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(381)
				p.Expression()
			}

			p.SetState(386)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(389)
		p.Match(StrictusParserT__4)
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(StrictusParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, StrictusParserRULE_eos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(391)
			p.Match(StrictusParserT__12)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(392)
			p.Match(StrictusParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(393)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "p.lineTerminatorAhead()", ""))
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(394)

		if !(p.GetTokenStream().LT(1).GetText() == "}") {
			panic(antlr.NewFailedPredicateException(p, "p.GetTokenStream().LT(1).GetText() == \"}\"", ""))
		}

	}

	return localctx
}

func (p *StrictusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *OrExpressionContext = nil
		if localctx != nil {
			t = localctx.(*OrExpressionContext)
		}
		return p.OrExpression_Sempred(t, predIndex)

	case 19:
		var t *AndExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AndExpressionContext)
		}
		return p.AndExpression_Sempred(t, predIndex)

	case 20:
		var t *EqualityExpressionContext = nil
		if localctx != nil {
			t = localctx.(*EqualityExpressionContext)
		}
		return p.EqualityExpression_Sempred(t, predIndex)

	case 21:
		var t *RelationalExpressionContext = nil
		if localctx != nil {
			t = localctx.(*RelationalExpressionContext)
		}
		return p.RelationalExpression_Sempred(t, predIndex)

	case 22:
		var t *AdditiveExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AdditiveExpressionContext)
		}
		return p.AdditiveExpression_Sempred(t, predIndex)

	case 23:
		var t *MultiplicativeExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MultiplicativeExpressionContext)
		}
		return p.MultiplicativeExpression_Sempred(t, predIndex)

	case 41:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *StrictusParser) OrExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) AndExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) EqualityExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) RelationalExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) AdditiveExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) MultiplicativeExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.lineTerminatorAhead()

	case 7:
		return p.GetTokenStream().LT(1).GetText() == "}"

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
