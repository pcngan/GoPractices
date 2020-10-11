package com.agidrive.mobileservice;

import java.util.HashSet;
import java.util.Set;

public class Solution2 {
    /**
     * This solution is check all the cases with poor performance if the String is large
     * We can use the Suffix tree to optimize the performance
     *
     * @param s
     * @return
     */
    public int solution(String s) {
        Set<String> checked = new HashSet<>();
        int result = s.length();
        for (int i = 1; i < s.length(); ++i) {
            checked.clear();
            for (int idx = 0; idx < s.length() - i; ++idx) {
                String substring = s.substring(idx, idx + i);
                if (!checked.contains(substring) && s.indexOf(substring, idx + 1) == -1) {
                    return i;
                }
                checked.add(substring);
            }
        }
        return result;
    }

    public static void main(String[] args) {
        Solution2 var = new Solution2();
        System.out.println(var.solution("abaaba"));
        System.out.println(var.solution("zyzyzyz"));
        System.out.println(var.solution("abbbabaaa"));
    }
}
