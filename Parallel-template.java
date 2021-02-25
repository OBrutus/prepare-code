import java.io.*;
import java.util.*;
import java.math.*;
import java.util.stream.*;

/***
 *  o/p stored in HashMap<Integer, String> 
 *      mapping test case number (Integer) to output StringBuilder
 */

/**
 * Author :     AV
 * At   :       @Home
 * Date :       21.2.
 * ver. :       0.0
 * link :       
 * file :       cf/
 */

// public class I_AM_AV{
class I_AM_AV{
    static private I_AM_AV_Template sc=new I_AM_AV_Template("pii.txt");
    static private StringBuilder sb=new StringBuilder();
    static public HashMap<Integer, StringBuilder> res = new HashMap<>();

    public static void main(String[] args) {
        final int TEST_CASES = sc.nextInt();
        // final int TEST_CASES = 1;
        Thread mainThread = Thread.currentThread();
        mainThread.setPriority(Thread.MIN_PRIORITY);
        System.out.println("Initially active count : "+mainThread.getThreadGroup().activeCount());
        for(int T=1; T<=TEST_CASES; T++){
            int n=sc.nextInt();
            int[] a=sc.input(n);
            // String s = sc.nextInt();

            Thread r = new I_AM_AV_TestCase(T, a, n);
            r.setPriority(Thread.MAX_PRIORITY);
            r.start();

            System.out.println("Active count : "+mainThread.getThreadGroup().activeCount());
        }
        Thread.yield();
        System.out.println("Active count : "+mainThread.getThreadGroup().activeCount());
        hook(); // Hooking Main Thread at end of all threads 
        for(int i=1; i<=TEST_CASES; i++){
            // ensure ith test cases are done 
            // and ith thread is running no more
	    // IF YOU FIND BETTER LOGIC THAN MENTIONED IN HOOK PLEASE DO MENTION
            sb.append( "Case #"+i+": "+res.get(i) +"\n" );
        }
        System.out.print(sb);
    }
    static synchronized void hook(){
        Thread mainThread = Thread.currentThread();
        Thread[] list = new Thread[mainThread.getThreadGroup().activeCount()];
        mainThread.getThreadGroup().enumerate(list);
        for(Thread t : list) {
            System.out.println(t.getName() +" " +t.getPriority());
            try {
                if(t!=mainThread){
                    t.join(); 
		    // all other threads are joined via Main thread
                }
            }
            catch(Exception e){}
        }
    }  
}

class I_AM_AV_TestCase extends Thread{
    int testCase;
    int[] a;
    int n;

    public I_AM_AV_TestCase(int testCase, int[] a, int n){
        this.testCase = testCase;
        this.a = a;
        this.n = n;
    }

    @Override
    public void run() {
        super.run();
        StringBuilder res = new StringBuilder();
        /*** LOGIC STARTS HERE */
        // --------------------

        /*** OUTPUT PROVIDING */
        I_AM_AV.res.put(testCase, res);
    }

}

class I_AM_AV_Template{ 
    static private BufferedReader br;
    static private StringTokenizer st; 
    static public final int MOD=1000000007; //10^9+7

    I_AM_AV_Template() {
        br = new BufferedReader(new InputStreamReader(System.in));
    }
    I_AM_AV_Template(String input) {
        try {
            // System.setIn(new FileInputStream(input));
            // br = new BufferedReader(new InputStreamReader(System.in));
            br = new BufferedReader(new InputStreamReader(new FileInputStream(input)));
        } 
        catch (Exception e) {
            e.printStackTrace();
        }
    }
    I_AM_AV_Template(String input, String output) {
        try {
            System.setIn(new FileInputStream(input));
            System.setOut(new PrintStream(output));
            br = new BufferedReader(new InputStreamReader(System.in));
        } 
        catch (Exception e) {
            e.printStackTrace();
        }
    }

    /**
     * Use StringBuilder for output
     */
    String next() { 
        while (st == null || !st.hasMoreElements()) { 
            try{ 
                st = new StringTokenizer(br.readLine()); 
            } 
            catch (IOException e) { 
                e.printStackTrace(); 
            } 
        } 
        return st.nextToken(); 
    } 
    int nextInt() { return Integer.parseInt(next()); } 
    long nextLong() { return Long.parseLong(next()); } 
    double nextDouble() {  return Double.parseDouble(next()); } 
    String nextLine() { 
        String str = ""; 
        try{ 
            str = br.readLine(); 
        } 
        catch (IOException e) { 
            e.printStackTrace(); 
        } 
        return str; 
    }
    public int[] input(int n){
        int[] a=new int[n];
        for(int i=0;i<n;i++)
            a[i]=this.nextInt();
        return a;
    }
    public int sum(int[] a, int start, int end){
        return Arrays.stream(a, start, end).sum();
    }
    public <E>void display(E[] a, int start, int end){
        for (;start<end; start++) {
            System.out.print(a[start]+" ");
        }
        System.out.println();
    }
    public void displayMatrix(int[][] a){
        for(int[] x: a)
            System.out.println(Arrays.toString(x));
    }

    public int max(int[] arr){ return Arrays.stream(arr).max().getAsInt(); }
    public int min(int[] arr){ return Arrays.stream(arr).min().getAsInt(); }

    public <E> E max(E[] a, int n){
        E max=a[0];
        return max;
    }
    public boolean isPrime(int number){
        if(number <= 2)
            return number == 2;
        else
            return  (number % 2) != 0
                    &&
                IntStream.rangeClosed(3, (int) Math.sqrt(number))
                    .filter(n -> n % 2 != 0)
                    .noneMatch(n -> (number % n == 0));
        }
}


