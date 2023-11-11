import java.io.*;
import java.util.*;
import java.math.*;
import java.util.stream.*;

/**
 * Author       :   {USERNAME}
 * Date         :   {DATE}
 * ver.         :   0.0
 * link         :   {PROBLEM_LINK}
 * dir          :   {PLATFORM}/
 */

// public class OBrutusSolution {
class OBrutusSolution {
    private static final OBrutusSolutionTemplate sc = new OBrutusSolutionTemplate();
    private static final StringBuilder sb = new StringBuilder();

    public static int code(int[] a, int n) {
        return 0;
    }

    public static void main(String[] args) {
        // final int TOTAL_TEST_CASES = 1;
        final int TOTAL_TEST_CASES = sc.nextInt();
        
        for (int testCase = TOTAL_TEST_CASES - 1; testCase >= 0; testCase--) {
            // Taking input
            // int n = sc.nextInt();
            // String s = sc.next();
            // int[] a = sc.input(n);

            // Compute
            // var res = code(a, n);
            // System.out.println("DEBUG: Case #" + (TOTAL_TEST_CASES - testCase) + ": " + res);

            // Storing the result
            // sb.append(res + "\n");
        }
        System.out.print(sb);
    }
}

class OBrutusSolutionTemplate {
    private static BufferedReader br;
    private static StringTokenizer st;
    public static final int MOD = 1_000_000_007; // 10^9+7

    OBrutusSolutionTemplate() {
        br = new BufferedReader(new InputStreamReader(System.in));
    }

    OBrutusSolutionTemplate(String input) {
        try {
            // System.setIn(new FileInputStream(input));
            // br = new BufferedReader(new InputStreamReader(System.in));
            br = new BufferedReader(new InputStreamReader(new FileInputStream(input)));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    OBrutusSolutionTemplate(String input, String output) {
        try {
            System.setIn(new FileInputStream(input));
            System.setOut(new PrintStream(output));
            br = new BufferedReader(new InputStreamReader(System.in));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    /**
     * Use StringBuilder for output
     */
    String next() {
        while (st == null || !st.hasMoreElements()) {
            try {
                st = new StringTokenizer(br.readLine());
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
        return st.nextToken();
    }

    int nextInt() {
        return Integer.parseInt(next());
    }

    long nextLong() {
        return Long.parseLong(next());
    }

    double nextDouble() {
        return Double.parseDouble(next());
    }

    String nextLine() {
        String str = "";
        try {
            str = br.readLine();
        } catch (IOException e) {
            e.printStackTrace();
        }
        return str;
    }

    public int[] input(int n) {
        int[] a = new int[n];
        for (int i = 0; i < n; i++)
            a[i] = this.nextInt();
        return a;
    }

    public int[] input(int[] a, int n) {
        for (int i = 0; i < n; i++)
            a[i] = this.nextInt();
        return a;
    }

    public int[][] input(int r, int c){
        int[][] a = new int[r][c];
        for(int[] x : a) {
            input(x, c);
        }
        return a;
    }

    public void displayMatrix(int[][] a) {
        for (int[] x : a)
            System.out.println(Arrays.toString(x));
    }

    public int gcd(int a, int b) {
        if(b==0) return a;
        return gcd(b, a%b);
    }

    public boolean isPrime(int number) {
        if (number <= 2)
            return number == 2;
        else
            return (number % 2) != 0 && IntStream.rangeClosed(3, (int) Math.sqrt(number)).filter(n -> n % 2 != 0)
                    .noneMatch(n -> (number % n == 0));
    }
}
