package main

import (
	"net/http"
	"log"
)

var favicon = []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 252, 0, 0, 0, 252, 8, 6, 0, 0, 0, 83, 171, 201, 103, 0, 0, 0, 4, 115, 66, 73, 84, 8, 8, 8, 8, 124, 8, 100, 136, 0, 0, 0, 9, 112, 72, 89, 115, 0, 0, 14, 133, 0, 0, 14, 133, 1, 184, 77, 169, 150, 0, 0, 0, 25, 116, 69, 88, 116, 83, 111, 102, 116, 119, 97, 114, 101, 0, 119, 119, 119, 46, 105, 110, 107, 115, 99, 97, 112, 101, 46, 111, 114, 103, 155, 238, 60, 26, 0, 0, 9, 231, 73, 68, 65, 84, 120, 156, 237, 221, 111, 172, 151, 101, 29, 199, 241, 55, 2, 130, 8, 254, 101, 19, 18, 169, 20, 204, 205, 153, 70, 130, 200, 159, 70, 206, 63, 171, 220, 178, 81, 154, 166, 46, 215, 170, 181, 181, 185, 53, 27, 79, 122, 64, 15, 91, 250, 184, 90, 205, 24, 211, 254, 79, 87, 61, 49, 209, 98, 2, 146, 56, 42, 214, 86, 9, 193, 74, 76, 96, 241, 55, 48, 242, 112, 128, 30, 92, 231, 180, 51, 6, 9, 231, 254, 253, 238, 239, 117, 159, 239, 251, 181, 93, 15, 239, 93, 159, 31, 231, 254, 112, 93, 231, 62, 247, 239, 190, 65, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 234, 140, 113, 209, 1, 70, 97, 34, 240, 94, 224, 226, 161, 33, 181, 237, 208, 208, 216, 1, 12, 6, 103, 57, 39, 93, 41, 252, 181, 192, 131, 192, 157, 192, 60, 74, 233, 165, 104, 199, 128, 205, 192, 243, 192, 211, 192, 214, 216, 56, 221, 247, 110, 96, 53, 112, 2, 56, 233, 112, 84, 62, 126, 9, 204, 69, 231, 108, 34, 240, 4, 48, 64, 252, 15, 209, 225, 56, 151, 49, 0, 60, 78, 165, 187, 208, 26, 183, 244, 151, 1, 63, 5, 110, 139, 14, 34, 53, 176, 14, 88, 14, 252, 51, 58, 200, 72, 181, 21, 126, 6, 229, 31, 106, 78, 116, 16, 169, 7, 182, 1, 75, 129, 61, 209, 65, 134, 157, 23, 29, 96, 132, 201, 192, 51, 88, 118, 141, 29, 115, 129, 103, 129, 73, 209, 65, 134, 141, 143, 14, 48, 194, 247, 128, 187, 163, 67, 72, 61, 118, 21, 48, 19, 248, 69, 116, 16, 168, 103, 75, 191, 20, 120, 41, 58, 132, 212, 71, 75, 129, 245, 209, 33, 106, 41, 252, 6, 96, 81, 116, 8, 169, 143, 94, 1, 110, 165, 92, 201, 15, 83, 195, 239, 240, 31, 195, 178, 107, 236, 187, 5, 184, 43, 58, 68, 13, 133, 127, 36, 58, 128, 212, 146, 240, 115, 61, 122, 75, 63, 1, 216, 15, 76, 11, 206, 33, 181, 225, 48, 229, 62, 147, 176, 251, 239, 163, 87, 248, 155, 176, 236, 202, 99, 26, 240, 254, 200, 0, 209, 133, 191, 46, 120, 126, 169, 109, 161, 231, 124, 116, 225, 103, 4, 207, 47, 181, 45, 244, 156, 159, 16, 57, 57, 48, 165, 225, 241, 63, 4, 94, 111, 112, 252, 138, 6, 199, 190, 62, 52, 255, 104, 45, 1, 22, 55, 56, 62, 243, 103, 111, 234, 27, 13, 142, 157, 13, 220, 223, 224, 248, 169, 13, 142, 237, 188, 149, 52, 251, 102, 210, 178, 134, 243, 55, 153, 251, 55, 13, 231, 94, 217, 112, 254, 101, 13, 231, 239, 242, 103, 111, 58, 154, 88, 214, 112, 238, 149, 13, 231, 111, 36, 122, 75, 47, 169, 69, 22, 94, 74, 196, 194, 75, 137, 88, 120, 41, 17, 11, 47, 37, 98, 225, 165, 68, 44, 188, 148, 136, 133, 151, 18, 177, 240, 82, 34, 22, 94, 74, 196, 194, 75, 137, 88, 120, 41, 17, 11, 47, 37, 98, 225, 165, 68, 162, 191, 15, 47, 69, 88, 211, 224, 216, 75, 123, 150, 34, 128, 133, 87, 70, 183, 71, 7, 136, 226, 150, 94, 74, 196, 194, 75, 137, 88, 120, 41, 17, 11, 47, 37, 98, 225, 165, 68, 44, 188, 148, 136, 133, 151, 18, 177, 240, 82, 34, 22, 94, 74, 196, 194, 75, 137, 88, 120, 41, 17, 11, 47, 37, 98, 225, 165, 68, 44, 188, 148, 72, 244, 215, 99, 7, 26, 30, 223, 244, 181, 197, 77, 44, 163, 249, 171, 135, 155, 200, 252, 217, 187, 236, 237, 200, 201, 163, 87, 248, 67, 193, 243, 75, 109, 59, 24, 57, 121, 116, 225, 119, 4, 207, 47, 181, 109, 123, 228, 228, 209, 133, 255, 3, 110, 13, 149, 199, 73, 96, 75, 100, 128, 232, 194, 239, 2, 254, 20, 156, 65, 106, 203, 31, 129, 61, 145, 1, 162, 11, 15, 240, 227, 232, 0, 82, 75, 126, 20, 29, 96, 92, 116, 0, 96, 6, 240, 87, 224, 194, 232, 32, 82, 31, 29, 1, 230, 224, 10, 207, 110, 224, 137, 232, 16, 82, 159, 125, 147, 224, 178, 67, 29, 43, 60, 192, 52, 96, 43, 101, 181, 151, 198, 154, 55, 129, 247, 81, 86, 249, 80, 53, 172, 240, 0, 135, 129, 79, 18, 124, 83, 130, 212, 7, 199, 128, 7, 168, 160, 236, 0, 227, 163, 3, 140, 176, 19, 120, 3, 184, 39, 58, 136, 212, 67, 95, 2, 158, 137, 14, 49, 172, 166, 194, 67, 249, 187, 252, 126, 224, 14, 234, 217, 125, 72, 163, 49, 8, 60, 10, 124, 59, 58, 200, 72, 181, 252, 14, 127, 170, 59, 40, 127, 174, 235, 244, 123, 188, 148, 214, 126, 224, 62, 224, 133, 232, 32, 167, 170, 109, 133, 31, 182, 3, 120, 10, 184, 24, 184, 17, 87, 123, 117, 195, 113, 224, 73, 224, 94, 130, 239, 168, 235, 178, 235, 129, 213, 148, 139, 30, 39, 29, 142, 10, 199, 97, 202, 57, 122, 61, 149, 171, 117, 75, 127, 58, 83, 128, 37, 192, 2, 224, 90, 96, 58, 48, 49, 52, 145, 178, 58, 6, 236, 165, 252, 41, 249, 21, 96, 61, 112, 52, 52, 145, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 52, 246, 116, 233, 251, 240, 80, 30, 103, 61, 143, 242, 64, 255, 233, 193, 89, 148, 219, 94, 202, 11, 84, 54, 83, 201, 19, 105, 207, 70, 23, 10, 63, 25, 120, 8, 120, 16, 88, 68, 252, 59, 237, 165, 145, 6, 129, 13, 148, 71, 178, 61, 5, 252, 39, 54, 78, 119, 157, 7, 124, 138, 242, 124, 187, 232, 71, 24, 57, 28, 103, 51, 118, 2, 95, 160, 222, 103, 69, 86, 187, 194, 207, 2, 158, 5, 110, 142, 14, 34, 141, 194, 38, 96, 57, 229, 61, 11, 85, 169, 177, 240, 11, 41, 101, 247, 181, 83, 234, 178, 93, 192, 39, 40, 207, 188, 171, 70, 109, 133, 95, 8, 252, 26, 184, 32, 58, 136, 212, 3, 71, 129, 101, 148, 21, 191, 10, 53, 21, 254, 93, 148, 127, 152, 43, 163, 131, 72, 61, 180, 27, 152, 79, 37, 219, 251, 90, 94, 240, 48, 30, 248, 57, 150, 93, 99, 207, 12, 224, 103, 84, 210, 181, 42, 66, 0, 15, 227, 5, 58, 141, 93, 183, 0, 159, 137, 14, 1, 117, 108, 233, 39, 3, 175, 1, 179, 163, 131, 72, 125, 244, 6, 229, 5, 42, 161, 47, 172, 168, 97, 133, 127, 8, 203, 174, 177, 111, 22, 229, 61, 241, 161, 106, 41, 188, 148, 65, 248, 182, 62, 122, 75, 63, 13, 216, 135, 239, 136, 83, 14, 3, 192, 229, 4, 222, 123, 31, 189, 194, 207, 195, 178, 43, 143, 243, 129, 15, 68, 6, 136, 46, 252, 220, 224, 249, 165, 182, 133, 158, 243, 209, 133, 191, 60, 120, 126, 169, 109, 161, 231, 124, 244, 87, 77, 39, 55, 60, 254, 49, 96, 75, 131, 227, 215, 52, 56, 118, 203, 208, 252, 163, 245, 48, 205, 46, 88, 102, 254, 236, 93, 54, 37, 114, 242, 232, 194, 55, 189, 104, 184, 25, 88, 219, 131, 28, 163, 113, 0, 120, 161, 193, 241, 75, 26, 206, 159, 249, 179, 119, 89, 232, 133, 242, 232, 45, 189, 164, 22, 89, 120, 41, 17, 11, 47, 37, 98, 225, 165, 68, 44, 188, 148, 136, 133, 151, 18, 177, 240, 82, 34, 22, 94, 74, 196, 194, 75, 137, 88, 120, 41, 17, 11, 47, 37, 98, 225, 165, 68, 44, 188, 148, 136, 133, 151, 18, 177, 240, 82, 34, 22, 94, 74, 196, 194, 75, 137, 88, 120, 41, 17, 11, 47, 37, 98, 225, 165, 68, 44, 188, 148, 136, 133, 151, 18, 177, 240, 82, 34, 22, 94, 74, 196, 194, 75, 137, 88, 120, 41, 17, 11, 47, 37, 98, 225, 165, 68, 44, 188, 148, 72, 244, 219, 99, 79, 54, 60, 254, 131, 196, 125, 134, 75, 129, 219, 27, 28, 127, 117, 195, 249, 51, 127, 246, 46, 107, 122, 206, 119, 218, 10, 202, 63, 128, 195, 145, 101, 124, 149, 64, 209, 91, 250, 125, 193, 243, 75, 109, 11, 61, 231, 163, 11, 191, 45, 120, 126, 169, 109, 161, 231, 252, 184, 200, 201, 129, 105, 148, 255, 241, 38, 6, 231, 144, 218, 48, 0, 92, 6, 188, 21, 21, 32, 122, 133, 63, 12, 108, 12, 206, 32, 181, 101, 61, 129, 101, 135, 248, 194, 3, 60, 21, 29, 64, 106, 201, 211, 209, 1, 162, 183, 244, 0, 147, 129, 215, 128, 217, 209, 65, 164, 62, 250, 27, 112, 29, 240, 118, 100, 136, 241, 145, 147, 15, 25, 4, 14, 0, 247, 68, 7, 145, 250, 232, 203, 192, 239, 163, 67, 212, 176, 194, 67, 249, 213, 98, 35, 176, 32, 58, 136, 212, 7, 27, 129, 37, 192, 137, 232, 32, 181, 20, 30, 96, 38, 240, 42, 112, 101, 116, 16, 169, 135, 118, 81, 22, 178, 55, 162, 131, 64, 29, 23, 237, 134, 237, 2, 150, 3, 71, 163, 131, 72, 61, 114, 20, 248, 56, 149, 148, 29, 234, 42, 60, 192, 43, 192, 98, 96, 103, 116, 16, 169, 161, 55, 129, 101, 148, 93, 107, 53, 106, 43, 60, 148, 11, 27, 183, 82, 202, 47, 117, 209, 70, 96, 62, 176, 41, 58, 200, 169, 106, 44, 60, 192, 63, 40, 165, 191, 23, 216, 17, 156, 69, 58, 91, 59, 129, 47, 2, 75, 41, 43, 124, 117, 106, 186, 104, 119, 38, 147, 129, 7, 128, 7, 41, 87, 58, 189, 13, 87, 53, 25, 160, 220, 65, 247, 244, 208, 8, 253, 59, 251, 59, 233, 66, 225, 71, 186, 16, 152, 7, 92, 3, 92, 65, 185, 23, 63, 250, 59, 253, 202, 101, 144, 114, 75, 248, 30, 96, 59, 240, 59, 130, 111, 151, 149, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 198, 168, 46, 125, 31, 126, 42, 229, 73, 34, 11, 128, 57, 192, 116, 252, 46, 188, 98, 12, 2, 123, 41, 47, 134, 220, 4, 172, 195, 239, 196, 247, 204, 13, 148, 39, 137, 252, 155, 248, 119, 123, 59, 28, 167, 27, 111, 81, 94, 153, 118, 3, 26, 181, 171, 128, 85, 192, 113, 226, 127, 160, 14, 199, 217, 140, 227, 192, 147, 192, 44, 116, 78, 238, 164, 188, 126, 42, 250, 7, 232, 112, 140, 102, 252, 11, 184, 155, 10, 213, 240, 110, 185, 83, 61, 10, 172, 6, 166, 68, 7, 145, 70, 105, 18, 229, 137, 203, 7, 169, 236, 81, 213, 181, 21, 254, 17, 224, 91, 212, 251, 248, 108, 233, 108, 157, 7, 124, 132, 242, 70, 165, 205, 193, 89, 254, 167, 166, 171, 244, 75, 128, 23, 129, 243, 163, 131, 72, 61, 116, 140, 242, 43, 234, 218, 224, 28, 64, 61, 133, 191, 8, 216, 74, 121, 244, 180, 52, 214, 236, 2, 174, 5, 142, 68, 7, 169, 101, 75, 255, 53, 224, 163, 209, 33, 164, 62, 153, 70, 89, 233, 215, 6, 231, 168, 98, 133, 159, 73, 185, 129, 225, 194, 232, 32, 82, 31, 29, 161, 220, 48, 182, 39, 50, 68, 13, 23, 199, 62, 143, 101, 215, 216, 55, 21, 248, 92, 116, 136, 26, 10, 127, 95, 116, 0, 169, 37, 159, 142, 14, 16, 189, 165, 159, 73, 165, 111, 217, 148, 250, 224, 36, 229, 156, 15, 219, 214, 71, 175, 240, 55, 5, 207, 47, 181, 105, 28, 112, 99, 100, 128, 232, 194, 95, 29, 60, 191, 212, 182, 107, 34, 39, 143, 46, 252, 197, 193, 243, 75, 109, 187, 36, 114, 242, 232, 194, 55, 189, 171, 238, 195, 148, 109, 210, 104, 71, 19, 107, 27, 206, 253, 245, 134, 243, 103, 254, 236, 93, 54, 41, 114, 242, 232, 194, 75, 106, 145, 133, 151, 18, 177, 240, 82, 34, 22, 94, 74, 196, 194, 75, 137, 88, 120, 41, 17, 11, 47, 37, 98, 225, 165, 68, 44, 188, 148, 136, 133, 151, 18, 177, 240, 82, 34, 22, 94, 74, 196, 194, 75, 137, 88, 120, 41, 17, 11, 47, 37, 98, 225, 165, 68, 44, 188, 148, 136, 133, 151, 18, 177, 240, 82, 34, 22, 94, 74, 196, 194, 75, 137, 88, 120, 41, 17, 11, 47, 37, 98, 225, 165, 68, 44, 188, 148, 136, 133, 151, 18, 177, 240, 82, 34, 22, 94, 74, 196, 194, 75, 137, 88, 120, 41, 145, 9, 193, 243, 159, 104, 120, 252, 253, 192, 45, 189, 8, 50, 10, 179, 129, 21, 13, 142, 95, 220, 112, 254, 204, 159, 189, 203, 154, 158, 243, 157, 246, 24, 112, 210, 225, 72, 52, 190, 66, 160, 232, 45, 253, 238, 224, 249, 165, 182, 133, 158, 243, 209, 133, 255, 75, 240, 252, 82, 219, 254, 28, 57, 249, 184, 200, 201, 41, 215, 16, 246, 1, 23, 5, 231, 144, 218, 112, 8, 152, 14, 12, 70, 5, 136, 94, 225, 7, 129, 231, 131, 51, 72, 109, 121, 142, 192, 178, 67, 124, 225, 1, 190, 31, 29, 64, 106, 201, 170, 232, 0, 209, 91, 250, 97, 235, 201, 253, 167, 26, 141, 125, 191, 5, 22, 81, 174, 212, 135, 169, 165, 240, 11, 129, 151, 169, 39, 143, 212, 107, 75, 128, 13, 209, 33, 106, 216, 210, 67, 249, 223, 111, 85, 116, 8, 169, 79, 190, 75, 5, 101, 135, 186, 86, 212, 243, 129, 53, 192, 135, 162, 131, 72, 61, 244, 50, 112, 27, 240, 118, 116, 16, 168, 171, 240, 0, 87, 0, 235, 128, 185, 209, 65, 164, 30, 216, 6, 44, 5, 246, 68, 7, 25, 86, 203, 150, 126, 216, 30, 202, 239, 243, 47, 70, 7, 145, 26, 122, 137, 114, 33, 186, 154, 178, 3, 140, 143, 14, 112, 26, 71, 129, 31, 0, 23, 0, 243, 169, 51, 163, 116, 38, 3, 192, 227, 192, 103, 129, 35, 177, 81, 186, 103, 54, 240, 29, 224, 56, 241, 95, 122, 112, 56, 254, 223, 56, 1, 252, 4, 184, 134, 138, 213, 246, 59, 252, 153, 204, 1, 30, 0, 238, 2, 110, 166, 92, 224, 147, 162, 13, 0, 175, 2, 191, 162, 236, 74, 183, 199, 198, 121, 103, 93, 41, 252, 72, 19, 128, 247, 0, 151, 12, 13, 169, 109, 7, 129, 3, 192, 223, 9, 190, 85, 86, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 146, 36, 73, 26, 195, 254, 11, 191, 18, 87, 68, 84, 173, 157, 198, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}

func handleFavicon(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-type", "image/png")
	log.Println("[FAVICON]")
	w.Write(favicon)
}