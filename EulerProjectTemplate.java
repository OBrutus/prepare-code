/**
 * For https://projecteuler.net/
 * */
import java.util.Scanner;


public class EulerProjectTemplate {
    public static final int PROBLEM_NUMBER = -1;

    static {
        if (PROBLEM_NUMBER == -1)
            throw new RuntimeException("Define the PROBLEM NUMBER");
        System.out.println("==============================================");
        System.out.println("| Problem link:                              |");
        System.out.println("| https://projecteuler.net/problem=" + PROBLEM_NUMBER);
        System.out.println("==============================================");
    }

    private static int code(int n) {
        int res = 0;
        return res;
    }

    public static void main(String[] args) {
        final Scanner sc = new Scanner(System.in);
        System.out.print("Number of testcases :: ");
        final int NUMBER_OF_TESTS = sc.nextInt();
        final long[] time = new long[NUMBER_OF_TESTS];
        final int[] res = new int[NUMBER_OF_TESTS];

        for (int i = 0; i < NUMBER_OF_TESTS; i++) {
            System.out.print("[" + (i + 1) + "] Inputs :: ");
            int n = sc.nextInt();
            long startTime = System.nanoTime();
            res[i] = code(n);
            long endTime = System.nanoTime();
            time[i] = endTime - startTime;
        }

        System.out.println("Output");
        System.out.println("result \t\t time(ns)");
        for(int i = 0; i < NUMBER_OF_TESTS; i++) {
            System.out.println(res[i] + " \t\t " + time[i]);
        }

        sc.close();
    }
}
