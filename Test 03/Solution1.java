package com.agidrive.mobileservice;

public class Solution1 {
    public int solution(int n, int k) {
        int counter = 0;
        int kCopy = k;
        for (int i = Math.min(n, kCopy); i > 0; i--, i = Math.min(i, kCopy)) {
            kCopy -= i;
            counter++;
        }
        if (kCopy > 0) {
            return -1;
        }
        return counter;
    }

    public static void main(String[] args) {
        Solution1 var = new Solution1();
        System.out.println(var.solution(5, 8));
        System.out.println(var.solution(4, 10));
        System.out.println(var.solution(1, 2));
        System.out.println(var.solution(10, 5));
        System.out.println(var.solution(10, 56));
    }
}
